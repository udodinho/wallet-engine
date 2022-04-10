package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/domain/wallet"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/helpers"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/response"
	"github.com/udodinho/golangProjects/wallet-engine/internals/ports"
	"log"
	"net/http"
	"time"
)

type HttpHandler struct {
	walletService ports.WalletService
}

func NewHttpHandler(walletService ports.WalletService) *HttpHandler {
	return &HttpHandler{walletService: walletService}
}

func (h *HttpHandler) CreateWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user = &wallet.User{}
		hashedPass, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			fmt.Println(err)
		}

		user.Reference = uuid.New().String()
		user.CreatedAt = time.Now().UTC()
		user.HashedSecretKey = string(hashedPass)

		// Binding the json
		if errs := helpers.Decode(ctx, &user); errs != nil {
			fmt.Println(errs)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errs})
			return
		}

		// Getting user by email
		userEmail, err := h.walletService.GetUserByEmail(user.Email)
		if err != nil {
			log.Println(err)
			return
		}

		// Checking to see if a user exists
		if len(userEmail) == 0 {
			userData, err := h.walletService.CreateWallet(user)

			if err != nil {
				log.Println(err)
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{"Wallet created successfully": userData})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User wallet already exists"})
			return
		}
	}
}

func (h *HttpHandler) CreditWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("reference")
		transaction := &wallet.Transaction{}
		transaction.UserID = userID
		transaction.TransactionRef = uuid.New().String()

		errs := helpers.Decode(ctx, &transaction)
		if errs != nil {
			response.JSON(ctx, "Could not get transaction", http.StatusInternalServerError, nil)
			return
		}

		if transaction.Amount < 1000 {
			response.JSON(ctx, "Sorry, deposit amount must be greater than N1000.00", http.StatusBadRequest, nil)
			return
		}

		var hashPass string
		var isUserActive bool

		userPass, err := h.walletService.CheckPassword(userID)
		if err != nil {
			response.JSON(ctx, "Could not get userID", http.StatusInternalServerError, nil)
			return
		}

		for _, user := range userPass {
			hashPass = user.HashedSecretKey
			isUserActive = user.IsActive
		}
		if confirmedPass := helpers.CompareHashPassword(transaction.Password, []byte(hashPass)); !confirmedPass {
			response.JSON(ctx, "Password does not match", http.StatusBadRequest, nil)
			return
		}

		acct := &wallet.Wallet{}
		userWallet := &wallet.User{}
		userWallet.IsActive = isUserActive

		if userWallet.IsActive == false {
			response.JSON(ctx, "User is not active", http.StatusBadRequest, nil)
			return
		}

		bal, err := h.walletService.GetBalance(userID)
		if err != nil {
			response.JSON(ctx, "Could not get user balance", http.StatusInternalServerError, nil)
			return
		}

		var currBal float64
		for _, userBal := range bal {
			currBal = userBal.Balance
		}
		acct.Balance = currBal

		acct.CreditWallet(transaction.Amount, userID)

		userTxn, err := h.walletService.SaveTransaction(transaction)
		if err != nil {
			response.JSON(ctx, "Could not save transaction", http.StatusInternalServerError, nil)
			return
		}

		currAcct, err := h.walletService.PostTransaction(acct)
		if err != nil {
			response.JSON(ctx, "Could not post transaction", http.StatusInternalServerError, nil)
			return
		}

		response.JSON(ctx, "Deposited successfully", http.StatusOK, gin.H{"transaction": userTxn, "balance": currAcct})

	}
}

func (h *HttpHandler) DebitWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("reference")
		transaction := &wallet.Transaction{}
		transaction.UserID = userID
		transaction.TransactionRef = uuid.New().String()

		errs := helpers.Decode(ctx, &transaction)
		if errs != nil {
			response.JSON(ctx, "Could not get transaction", http.StatusInternalServerError, nil)
			return
		}

		if transaction.Amount < 1000 {
			response.JSON(ctx, "Sorry, withdrawal amount must be greater than N1000.00", http.StatusBadRequest, nil)
			return
		}

		var hashPass string
		var isUserActive bool

		userPass, err := h.walletService.CheckPassword(userID)
		if err != nil {
			response.JSON(ctx, "Could not get userID", http.StatusInternalServerError, nil)
			return
		}

		for _, user := range userPass {
			hashPass = user.HashedSecretKey
			isUserActive = user.IsActive
		}

		if confirmedPass := helpers.CompareHashPassword(transaction.Password, []byte(hashPass)); !confirmedPass {
			response.JSON(ctx, "Password does not match", http.StatusBadRequest, nil)
			return
		}

		acct := &wallet.Wallet{}
		userWallet := &wallet.User{}
		userWallet.IsActive = isUserActive

		if userWallet.IsActive == false {
			response.JSON(ctx, "User is not active", http.StatusBadRequest, nil)
			return
		}

		bal, err := h.walletService.GetBalance(userID)
		if err != nil {
			response.JSON(ctx, "Could not get user balance", http.StatusInternalServerError, nil)
			return
		}

		var currBal float64
		for _, userBal := range bal {
			currBal = userBal.Balance
		}
		acct.Balance = currBal

		if acct.Balance <= 0 {
			response.JSON(ctx, "Insufficient funds", http.StatusBadRequest, nil)
			return
		}

		if acct.Balance < transaction.Amount {
			response.JSON(ctx, "Sorry, you don't have sufficient funds for this transaction", http.StatusBadRequest, nil)
			return
		}

		acct.DebitWallet(transaction.Amount, userID)

		userTxn, err := h.walletService.SaveTransaction(transaction)
		if err != nil {
			response.JSON(ctx, "Could not save transaction", http.StatusInternalServerError, nil)
			return
		}

		currAcct, err := h.walletService.PostTransaction(acct)
		if err != nil {
			response.JSON(ctx, "Debit was not successful", http.StatusInternalServerError, nil)
			return
		}

		response.JSON(ctx, "Your account was debited successfully", http.StatusOK, gin.H{"transaction": userTxn, "balance": currAcct})

	}
}

func (h *HttpHandler) ActivateDeactivate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRef := ctx.Param("reference")
		activate := ctx.Query("activate")

		user := &wallet.User{}

		var msg string
		var status bool

		if activate == "true" {
			status = true
			msg = "User activated successfully"
		} else {
			status = false
			msg = "User deactivated successfully"
		}

		user.ActivateWallet(status)
		_, err := h.walletService.ChangeStatus(user.IsActive, userRef)
		if err != nil {
			response.JSON(ctx, "Could not change status", http.StatusInternalServerError, nil)
			return
		}

		response.JSON(ctx, msg, http.StatusOK, nil)
	}
}
