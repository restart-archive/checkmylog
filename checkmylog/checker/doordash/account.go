package doordash

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
		if len(pm.Type) <= 0 && len(pm.Last4) <= 0 {
			continue
		}
		pms = append(pms, fmt.Sprintf("[%s %s]", pm.Type, pm.Last4))
	}
	return pms
}

// FileName ...
func (a account) FileName() string {
	return fmt.Sprintf("DOORDASH %dPM - %s", len(a.PaymentMethods()), time.Now().Format(time.RFC3339))
}
