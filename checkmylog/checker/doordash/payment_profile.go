package doordash

import "time"

// account represents an account with payment profiles.
type paymentProfile struct {
	Typename                        string `json:"__typename"`
	ID                              string `json:"id"`
	IsDefault                       bool   `json:"isDefault"`
	Type                            string `json:"type"`
	StripeID                        string `json:"stripeId"`
	Last4                           string `json:"last4"`
	ExpYear                         string `json:"expYear"`
	ExpMonth                        string `json:"expMonth"`
	HasExistingCard                 any    `json:"hasExistingCard"`
	CardBenefitMembershipLinkStatus any    `json:"cardBenefitMembershipLinkStatus"`
	Metadata                        struct {
		Typename                  string `json:"__typename"`
		PaypalAccount             string `json:"paypalAccount"`
		SetupIntentClientSecret   string `json:"setupIntentClientSecret"`
		SetupIntentRequiresAction bool   `json:"setupIntentRequiresAction"`
		IsDashCard                bool   `json:"isDashCard"`
		IsHsaFsaCard              bool   `json:"isHsaFsaCard"`
	} `json:"metadata"`
	AccountDisplayDescription string `json:"accountDisplayDescription"`
	PaymentMethodUUID         string `json:"paymentMethodUuid"`
	PaymentMethodAvailability struct {
		Typename                        string `json:"__typename"`
		PaymentMethodAvailabilityStatus string `json:"paymentMethodAvailabilityStatus"`
	} `json:"paymentMethodAvailability"`
	CreatedAt           time.Time `json:"createdAt"`
	IsPrimaryCardHolder any       `json:"isPrimaryCardHolder"`
}
