package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/domain/wallet"
	mock "github.com/udodinho/golangProjects/wallet-engine/internals/core/services/mock"
	"go/types"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockWalletService(ctrl)
	router := gin.Default()
	handler := NewHttpHandler(mockService)
	router.POST("/api/v1/createWallet", handler.CreateWallet())

	newUser := &wallet.User{
		Reference: "123",
		FirstName: "Ruth",
		LastName:  "Mendes",
		Email:     "rmendes@g.com",
		Password:  "password",
	}

	jsn, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("Error marshalling user: %v", err)
	}

	var user []*wallet.User

	mockService.EXPECT().GetUserByEmail(newUser.Email).Return(user, nil)
	mockService.EXPECT().CreateWallet(gomock.Any()).Return(newUser, nil)

	t.Run("Test create wallet", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/v1/createWallet", bytes.NewBuffer(jsn))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		log.Println(resp.Body)
		if resp.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.Code)
			return
		}
	})
}

func TestCreditWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockWalletService(ctrl)
	router := gin.Default()
	handler := NewHttpHandler(mockService)
	router.POST("/api/v1/creditWallet/:reference", handler.CreditWallet())

	trnx := &wallet.Transaction{
		UserID:         "oiuyfdfcv",
		TransactionRef: "wscvhtfdx",
		PhoneNumber:    "09098765432",
		Amount:         19000.00,
		Password:       "pass123345",
	}

	jsn, err := json.Marshal(trnx)
	if err != nil {
		t.Errorf("Error marshalling user: %v", err)
	}

	var newUser []*wallet.User
	mockService.EXPECT().CheckPassword(gomock.Any()).Return(newUser, nil)

	t.Run("test for credit wallet", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/v1/creditWallet/:reference", bytes.NewBuffer(jsn))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		log.Println(resp.Body)
		assert.Contains(t, resp.Body.String(), "User is not active")
	})
}

func TestDebitWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockWalletService(ctrl)
	router := gin.Default()
	handler := NewHttpHandler(mockService)
	router.POST("/api/v1/debitWallet/:reference", handler.DebitWallet())

	trnx := &wallet.Transaction{
		UserID:         "oiuyfdfcv",
		TransactionRef: "wscvhtfdx",
		PhoneNumber:    "09098765432",
		Amount:         19000.00,
		Password:       "pass123345",
	}

	jsn, err := json.Marshal(trnx)
	if err != nil {
		t.Errorf("Error marshalling user: %v", err)
	}

	var newUser []*wallet.User

	mockService.EXPECT().CheckPassword(gomock.Any()).Return(newUser, nil)

	t.Run("test for debit wallet", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/v1/debitWallet/:reference", bytes.NewBuffer(jsn))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		log.Println(resp.Body)
		assert.Contains(t, resp.Body.String(), "User is not active")
	})
}

func TestActivateDeactivate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockWalletService(ctrl)
	router := gin.Default()
	handler := NewHttpHandler(mockService)
	router.PUT("/api/v1/activateDeactivateWallet/:reference", handler.ActivateDeactivate())

	newUser := &wallet.User{
		Reference:   "2f509a97-12ac-7cc9-baf0-fd01a5d653a0",
		FirstName:   "King",
		LastName:    "Daniel",
		Email:       "kdaniel@g.com",
		Password:    "password",
		DateOfBirth: "01/01/1990",
		IsActive:    true,
	}

	stats := []byte("status changed")

	t.Run("Test for activate-deactivate wallet", func(t *testing.T) {
		mockService.EXPECT().ChangeStatus(gomock.Any(), newUser.Reference).Return(types.Interface{}, nil)
		req, err := http.NewRequest(http.MethodPut, "/api/v1/activateDeactivateWallet/2f509a97-12ac-7cc9-baf0-fd01a5d653a0", bytes.NewBuffer(stats))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		log.Println("here", resp.Body)
		assert.Contains(t, resp.Body.String(), "User deactivated successfully")
		assert.Equal(t, http.StatusOK, resp.Code)

	})

}
