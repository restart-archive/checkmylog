package uber

// accountDetails is the account details.
type accountDetails struct {
	AccountDetails struct {
		Loading   bool        `json:"loading"`
		Data      interface{} `json:"data"`
		Error     interface{} `json:"error"`
		ProductID interface{} `json:"productId"`
	} `json:"accountDetails"`
	AccountFeed struct {
		Data      interface{} `json:"data"`
		Error     interface{} `json:"error"`
		Loading   bool        `json:"loading"`
		ProductID interface{} `json:"productId"`
	} `json:"accountFeed"`
	PaymentProfiles struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"paymentProfiles"`
	PaymentPreferences struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"paymentPreferences"`
	Profiles struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"profiles"`
	TransactionDetails struct {
		CurrentTransactionID              string `json:"currentTransactionId"`
		IsTransactionDetailsPanelExpanded bool   `json:"isTransactionDetailsPanelExpanded"`
	} `json:"transactionDetails"`
	UberCashBalance struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"uberCashBalance"`
	User struct {
		UUID    string      `json:"uuid"`
		Loading bool        `json:"loading"`
		Error   interface{} `json:"error"`
		Data    struct {
			UUID                 string `json:"uuid"`
			Firstname            string `json:"firstname"`
			Lastname             string `json:"lastname"`
			FullPictureURL       string `json:"fullPictureUrl"`
			FullName             string `json:"fullName"`
			Role                 string `json:"role"`
			ShowPostMatesMessage bool   `json:"showPostMatesMessage"`
		} `json:"data"`
	} `json:"user"`
	WalletHome struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"walletHome"`
	GiftRedeem struct {
		IsVideoRedemptionModalOpened bool `json:"isVideoRedemptionModalOpened"`
	} `json:"giftRedeem"`
	Statements struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"statements"`
	UberCashPurchase struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"uberCashPurchase"`
	AddFundsOptions struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"addFundsOptions"`
	HydrationData struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"hydrationData"`
	VerificationStatus struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"verificationStatus"`
	FinancialAccounts struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"financialAccounts"`
	InformationPage struct {
		Loading bool        `json:"loading"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	} `json:"informationPage"`
	NextStep struct {
		Data struct {
			NextScreenID string `json:"nextScreenID"`
		} `json:"data"`
		Error   interface{} `json:"error"`
		Loading bool        `json:"loading"`
	} `json:"nextStep"`
}
