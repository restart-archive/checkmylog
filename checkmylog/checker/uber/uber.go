package uber

import (
	"regexp"
)

var (
	// accountDetailsRegex is a regex to extract account details.
	accountDetailsRegex = regexp.MustCompile(`<script type="application/json" id="__REDUX_STATE__"> (.*?) </script>`)
	// baseURL is the base URL for the Uber wallet.
	baseURL = "https://wallet.uber.com/"
	// accountDetailsURL is the URL to get account details.
	paymentProfilesURL = baseURL + "api/getPaymentProfiles?localeCode=en-CA"
)
