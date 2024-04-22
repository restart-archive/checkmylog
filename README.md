## DISCLAIMER
This is a personal project and is not intended to use in a harmful way. I am not responsible for any damage caused by this project. Use it at your own risk.

## Description
This project is a high level Go API for making or / checking cookie contents.

## Usage
### Run:
`go get github.com/restartfu/checkmylog/checkmylog`
### You may then simply copy and paste the following code to your project. (main.go)
```go
package main

import (
	"bufio"
	"fmt"
	"github.com/restartfu/checkmylog/checkmylog"
	"github.com/restartfu/checkmylog/checkmylog/checker/doordash"
	"github.com/restartfu/checkmylog/checkmylog/checker/microsoft"
	"github.com/restartfu/checkmylog/checkmylog/checker/uber"
	"os"
	"strings"
)

func main() {
	var (
		filepath string
		scanner  = bufio.NewScanner(os.Stdin)
	)

	fmt.Print("Enter the path to the logs: ")
	for scanner.Scan() {
		filepath = strings.ReplaceAll(scanner.Text(), "\"", "")
		break
	}

	var checkers []checkmylog.Checker
	checkers = append(checkers, doordash.DoorDash{})
	checkers = append(checkers, uber.Uber{})
	checkers = append(checkers, microsoft.Microsoft{})

	attempts, success := checkmylog.CheckLogs(filepath, checkers...)

	fmt.Printf("\nAttempts: %d\nSuccess: %d\n", attempts, success)
}
```

# Implementing your own checker
### Define a type which implements the`checkmylog.Checker` interface. It consists of the following:
```go
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
```
