package checkmylog

import (
	"fmt"
	tls "github.com/refraction-networking/utls"
	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/fhttp/cookiejar"
	"github.com/saucesteals/mimic"
	"net/url"
	"os"
)

// Log is a struct to represent a log.
type Log struct {
	client    *http.Client
	transport *http.Transport

	jar     *cookiejar.Jar
	content []byte
}

// mimicVersion is the latest version of the mimic browser.
var mimicVersion = mimic.MustGetLatestVersion(mimic.PlatformWindows)

// NewLog creates a new log with the specified cookie jar and content.
func NewLog(jar *cookiejar.Jar, content []byte) Log {
	m, err := mimic.Chromium(mimic.BrandChrome, mimicVersion)
	if err != nil {
		panic(err)
	}

	transport := m.ConfigureTransport(&http.Transport{
		TLSClientConfig: &tls.Config{},
	})

	client := &http.Client{
		Jar:       jar,
		Transport: transport,
	}

	l := Log{
		content:   content,
		client:    client,
		jar:       jar,
		transport: transport,
	}

	return l
}

func (l Log) saveToFile(c Checker, results Account) {
	_ = os.MkdirAll(fmt.Sprintf("%s/%s", outputPath, c.Name()), os.ModePerm)
	f, err := os.Create(fmt.Sprintf("%s/%s/%s.txt", outputPath, c.Name(), results.FileName()))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = f.Write(l.content)
}

// Cookies returns the cookies for the specified URL.
func (l Log) Cookies(u *url.URL) []*http.Cookie {
	return l.jar.Cookies(u)
}

// Do sends the specified request and returns the response.
func (l Log) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	return l.client.Do(req)
}
