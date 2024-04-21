package testing

import (
	"github.com/restartfu/checkmylog/checkmylog"
	"github.com/restartfu/checkmylog/checkmylog/checker/uber"
	"os"
	"strings"
	"testing"
)

func TestUber(t *testing.T) {
	filename := "./../assets/cookies/uber/2024-03-25 08-27-33_1.txt"
	content, _ := os.ReadFile(filename)

	jar, ok := checkmylog.LoadCookies(string(content))
	if !ok {
		panic("no cookie")
	}

	log := checkmylog.NewLog(jar, content)

	results, ok := (uber.Uber{}).Capture(log)
	if !ok {
		panic("no result")
	}

	panic(strings.Join(results.PaymentMethods(), ", "))
}
