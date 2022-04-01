package services

import (
	"github.com/udodinho/golangProjects/wallet-engine/internals/ports"
	"github.com/udodinho/goprojects/wallet-engine/domain/wallet"
)

type WalletService struct {
	WalletRepo ports.WalletRepository
}

func (w *WalletService) CreateWallet(wallet *wallet.User) (*wallet.User, error) {
	return w.WalletRepo.CreateWallet(wallet)
}

func (w *WalletService) GetUserByEmail(email string) ([]*wallet.User, error) {
	return w.WalletRepo.GetUserByEmail(email)
}

func (w *WalletService) CheckPassword(userRef string) ([]*wallet.User, error) {
	return w.WalletRepo.CheckPassword(userRef)
}

func (w *WalletService) GetAccountBalance(userID string) ([]*wallet.User, error) {
	return w.WalletRepo.GetAccountBalance(userID)
}

func (w *WalletService) ChangeStatus(isActive bool, userRef string) (interface{}, error) {
	return w.WalletRepo.ChangeStatus(isActive, userRef)
}

func (w *WalletService) PostTransaction(transaction *wallet.Wallet) (*wallet.Wallet, error) {
	return w.WalletRepo.PostTransaction(transaction)
}

func (w *WalletService) SaveTransaction(transaction *wallet.Wallet) (*wallet.Wallet, error) {
	return w.WalletRepo.SaveTransaction(transaction)
}

func New(wallet ports.WalletRepository) *WalletService {
	return &WalletService{
		WalletRepo: wallet,
	}
}
