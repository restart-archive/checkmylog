package uber

import "fmt"

// account represents an account with payment methods.
type account struct {
	// details is the account details.
	details accountDetails
	// paymentProfiles is the payment profiles.
	paymentProfiles []paymentProfile
}

// PaymentMethods ...
func (a account) PaymentMethods() []string {
	var pms []string
	for _, pm := range a.paymentProfiles {
		if len(pm.CardType) <= 0 && len(pm.CardNumber) <= 0 {
			continue
		}
		pms = append(pms, fmt.Sprintf("[%s %s]", pm.CardType, pm.CardNumber))
	}
	return pms
}

// FileName ...
func (a account) FileName() string {
	return fmt.Sprintf("UBER %dPM - %s", len(a.PaymentMethods()), a.details.User.Data.FullName)
}
