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
