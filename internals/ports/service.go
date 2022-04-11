package ports

import "github.com/udodinho/golangProjects/wallet-engine/internals/core/domain/wallet"

type WalletService interface {
	CreateWallet(wallet *wallet.User) (*wallet.User, error)
	GetUserByEmail(email string) ([]*wallet.User, error)
	CheckPassword(userRef string) ([]*wallet.User, error)
	GetBalance(userID string) ([]*wallet.Wallet, error)
	ChangeStatus(isActive bool, userRef string) (interface{}, error)
	PostTransaction(transaction *wallet.Wallet) (interface{}, error)
	SaveTransaction(transaction *wallet.Transaction) (interface{}, error)
}
