package wallet

// Transaction struct
type Transaction struct {
	UserID         string  `json:"user_id"bson:"user_id"`
	TransactionRef string  `json:"transaction_ref"bson:"transaction_ref"`
	PhoneNumber    string  `json:"phone_number"bson:"phone_number"`
	Amount         float64 `json:"amount"bson:"amount"`
	Password       string  `json:"-"bson:"password"`
}

// Wallet struct
type Wallet struct {
	UserID  string  `json:"user_id"bson:"user_id"`
	Balance float64 `json:"balance"bson:"balance"`
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

func (u *User) ActivateWallet(active bool) {
	u.IsActive = active
}
