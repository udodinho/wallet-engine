package domain

// Transaction struct
type Transaction struct {
	UserID         string
	TransactionRef string
	PhoneNumber    string
	Amount         float64
	Password       string
}

// Wallet struct
type Wallet struct {
	UserID    string
	Balance   float64
	Reference string
}

// CreditWallet credits the user wallet
func (w *Wallet) CreditWallet(money float64, userID string) {
	w.UserID = userID
	w.Balance = w.Balance + money
}

// DebitWallet debits user wallet
func (w *Wallet) DebitWallet(money float64, userID string) {
	w.UserID = userID
	w.Balance = w.Balance - money
}
