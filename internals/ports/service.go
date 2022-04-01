package ports

import "github.com/udodinho/goprojects/wallet-engine/domain/wallet"

type WalletService interface {
	CreateWallet(wallet *wallet.User) (*wallet.User, error)
	GetUserByEmail(email string) ([]*wallet.User, error)
	CheckPassword(userRef string) ([]*wallet.User, error)
	GetAccountBalance(userID string) ([]*wallet.User, error)
	ChangeStatus(isActive bool, userRef string) (interface{}, error)
	PostTransaction(transaction *wallet.Wallet) (*wallet.Wallet, error)
	SaveTransaction(transaction *wallet.Wallet) (*wallet.Wallet, error)
}
