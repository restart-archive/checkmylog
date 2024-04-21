package microsoft

import (
	"fmt"
	"time"
)

// account represents an account with payment profiles.
type account struct {
	// paymentProfiles is the payment profiles.
	paymentProfiles []paymentProfile
}

// PaymentMethods ...
func (a account) PaymentMethods() []string {
	var pms []string
	for _, pm := range a.paymentProfiles {
		if len(pm.Company()) <= 0 && len(pm.Digits()) <= 0 {
			continue
		}
		pms = append(pms, fmt.Sprintf("[%s %s]", pm.Company(), pm.Digits()))
	}
	return pms
}

// FileName ...
func (a account) FileName() string {
	return fmt.Sprintf("MICROSOFT %dPM - %s", len(a.PaymentMethods()), time.Now().Format(time.RFC3339))
}
