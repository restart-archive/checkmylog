package uber

import "time"

// paymentData represents payment data.
type paymentData struct {
	// Data is the payment data.
	Data struct {
		// PaymentProfiles is the payment profiles.
		PaymentProfiles []paymentProfile `json:"paymentProfiles"`
	} `json:"data"`
}

// paymentProfile represents a payment profile.
type paymentProfile struct {
	// AccountName is the account name.
	AccountName string `json:"accountName"`

	// Analytics is the analytics.
	Analytics struct {
		// PaymentMethodID is the payment method ID.
		PaymentMethodID string `json:"paymentMethodID"`
	} `json:"analytics"`

	// AuthenticationType is the authentication type.
	AuthenticationType string `json:"authenticationType"`
	// BillingCountryIso2 is the billing country ISO2.
	BillingCountryIso2 string `json:"billingCountryIso2"`
	// CardNumber is the card number.
	CardNumber string `json:"cardNumber"`
	// CardType is the card type.
	CardType string `json:"cardType"`
	// ClientUuid is the client UUID.
	ClientUUID string `json:"clientUuid"`
	// Displayable is the displayable.

	Displayable struct {
		// DisplayName is the display name.
		DisplayName string `json:"displayName"`
		// IconURL is the icon URL.
		IconURL string `json:"iconURL"`
	} `json:"displayable"`

	// HasBalance is the has balance.
	HasBalance bool `json:"hasBalance"`
	// Status is the status.
	Status string `json:"status"`
	// SupportedCapabilities is the supported capabilities.
	SupportedCapabilities []string `json:"supportedCapabilities"`
	// TokenDisplayName is the token display name.
	TokenDisplayName string `json:"tokenDisplayName"`
	// TokenType is the token type.
	TokenType string `json:"tokenType"`
	// UpdatedAt is the updated at.
	UpdatedAt time.Time `json:"updatedAt"`
	// UseCase is the use case.
	UseCase string `json:"useCase"`
	// UUID is the UUID.
	UUID string `json:"uuid"`
}
