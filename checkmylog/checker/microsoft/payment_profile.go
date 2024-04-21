package microsoft

// paymentProfile represents a payment profile.
type paymentProfile struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PaymentMethod struct {
		PaymentMethodType string `json:"paymentMethodType"`
		Properties        struct {
			OfflineRecurring         bool     `json:"offlineRecurring"`
			UserManaged              bool     `json:"userManaged"`
			SoldToAddressRequired    bool     `json:"soldToAddressRequired"`
			SplitPaymentSupported    bool     `json:"splitPaymentSupported"`
			SupportedOperations      []string `json:"supportedOperations"`
			Taxable                  bool     `json:"taxable"`
			ProviderRemittable       bool     `json:"providerRemittable"`
			IsNonStoredPaymentMethod bool     `json:"isNonStoredPaymentMethod"`
		} `json:"properties"`
		PaymentMethodGroup  string   `json:"paymentMethodGroup"`
		GroupDisplayName    string   `json:"groupDisplayName"`
		ExclusionTags       []string `json:"exclusionTags"`
		PaymentMethodFamily string   `json:"paymentMethodFamily"`
		Display             struct {
			Name  string `json:"name"`
			Logo  string `json:"logo"`
			Logos []struct {
				MimeType string `json:"mimeType"`
				URL      string `json:"url"`
			} `json:"logos"`
		} `json:"display"`
	} `json:"paymentMethod,omitempty"`
	Status              string `json:"status"`
	CreationDateTime    string `json:"creationDateTime,omitempty"`
	LastUpdatedDateTime string `json:"lastUpdatedDateTime,omitempty"`
	Details             struct {
		Exportable                       bool    `json:"exportable"`
		IsIndiaExpiryGroupDeleteFlighted bool    `json:"isIndiaExpiryGroupDeleteFlighted"`
		Email                            string  `json:"email"`
		BillingAgreementID               string  `json:"billingAgreementId"`
		FirstName                        string  `json:"firstName"`
		LastName                         string  `json:"lastName"`
		BillingAgreementType             string  `json:"billingAgreementType"`
		PicvRequired                     bool    `json:"picvRequired"`
		Balance                          float64 `json:"balance"`
		DefaultDisplayName               string  `json:"defaultDisplayName"`
		CardType                         string  `json:"cardType"`
		LastFourDigits                   string  `json:"lastFourDigits"`
		ExpiryYear                       string  `json:"expiryYear"`
		ExpiryMonth                      string  `json:"expiryMonth"`
		Lots                             []struct {
			CurrentBalance  float64 `json:"currentBalance"`
			OriginalBalance float64 `json:"originalBalance"`
			PendingBalance  float64 `json:"pendingBalance"`
			FundOrderID     int64   `json:"fundOrderId"`
			Type            string  `json:"type"`
			Status          string  `json:"status"`
			TokenInstanceID string  `json:"tokenInstanceId"`
			LastUpdatedTime string  `json:"lastUpdatedTime"`
		} `json:"lots"`
	} `json:"details,omitempty"`
}

// Company returns the company of the payment profile.
func (p paymentProfile) Company() string {
	if len(p.Details.Email) > 0 {
		return "Paypal"
	}
	return p.Details.CardType
}

// Digits returns the digits of the payment profile.
func (p paymentProfile) Digits() string {
	email := p.Details.Email

	if len(email) > 0 {
		return email
	}

	return p.Details.LastFourDigits
}
