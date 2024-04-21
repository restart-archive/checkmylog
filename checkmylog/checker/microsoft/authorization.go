package microsoft

import (
	"encoding/json"
	"github.com/restartfu/checkmylog/checkmylog"
	http "github.com/saucesteals/fhttp"
	"io"
	"time"
)

type Authorization struct {
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"tokenExpiry"`
	TokenScopes []string  `json:"tokenScopes"`
	IsSuccess   bool      `json:"isSuccess"`
}

func (m Microsoft) authorization(log checkmylog.Log) (string, bool) {
	var auths []Authorization
	req, err := http.NewRequest("GET", "https://account.microsoft.com/auth/acquire-onbehalf-of-token?scopes=pidl", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Referer", "https://account.microsoft.com/billing/payments?fref=home.drawers.payment-options.manage-payment&refd=account.microsoft.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Correlation-Context", "v=1,ms.b.tel.market=fr-FR,ms.b.tel.scenario=ust.amc.billing.payment-north-star,ms.c.ust.scenarioStep=PaymentNorthStarOboAuthStart")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	resp, err := log.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "", false
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(bodyText, &auths)
	if len(auths) <= 0 {
		return "", false
	}

	tok := auths[0].Token
	return tok, len(tok) > 0
}
