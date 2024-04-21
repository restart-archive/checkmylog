package checkmylog

// Account represents an account with payment methods.
type Account interface {
	// PaymentMethods returns a list of payment methods.
	PaymentMethods() []string
	// FileName returns a file name for the account.
	FileName() string
}
