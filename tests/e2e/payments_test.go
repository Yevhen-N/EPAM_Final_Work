package e2e

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

func TestUser(t *testing.T) {

	// User POST test
	createUserBody := apiv1.UserRequest{
		FullName: "barak obama",
		Email:    "usa@gmail.com",
		Password: "qwerty12345",
	}
	req, err := makeRequest(http.MethodPost, basePath+"/users", &createUserBody)
	require.NoError(t, err)

	var createUserResp apiv1.UserResponse

	err = doRequest(req, &createUserResp)
	require.NoError(t, err)

	require.NotNil(t, createUserResp)
	assert.Equal(t, createUserBody.FullName, createUserResp.FullName)
	assert.Equal(t, createUserBody.Email, createUserResp.Email)
	assert.NotEmpty(t, createUserResp.ID)
	assert.NotEmpty(t, createUserResp.Role)
	assert.NotEmpty(t, createUserResp.Status)
	assert.Equal(t, createUserResp.Status, model.UserStatusActive)
	assert.Equal(t, createUserResp.Role, model.UserRoleUser)

	// User PUT test
	updateAccountBody := apiv1.UserStatusRequest{
		ID:     createUserResp.ID,
		Status: model.UserStatusBlocked,
	}
	req, err = makeRequest(http.MethodPut, basePath+"/users/lock", &updateAccountBody)
	require.NoError(t, err)

	var updateUserResp apiv1.UserResponse

	err = doRequest(req, &updateUserResp)
	require.NoError(t, err)

	require.NotNil(t, updateUserResp)
	assert.Equal(t, updateUserResp.Status, model.UserStatusBlocked)
	assert.NotEmpty(t, updateUserResp.ID)

	// User GET test
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	var getUserResp apiv1.UserResponse

	err = doRequest(req, &getUserResp)
	require.NoError(t, err)

	require.NotNil(t, getUserResp)
	assert.NotEmpty(t, getUserResp.ID)
	assert.NotEmpty(t, getUserResp.Email)
	assert.NotEmpty(t, getUserResp.Status)
	assert.NotEmpty(t, getUserResp.Role)
	assert.NotEmpty(t, getUserResp.FullName)

	// Delete test user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, &apiv1.UserResponse{})
	require.Error(t, err)
}

func TestAccount(t *testing.T) {

	// User TEST POST
	createUserBody := apiv1.UserRequest{
		FullName: "bonapard napoleon",
		Email:    "fra@gmail.com",
		Password: "qwerty12345",
	}

	req, err := makeRequest(http.MethodPost, basePath+"/users", &createUserBody)
	require.NoError(t, err)

	var createUserResp apiv1.UserResponse

	err = doRequest(req, &createUserResp)
	require.NoError(t, err)

	// Account POST test
	createAccountBody := apiv1.AccountRequest{
		UserID:   createUserResp.ID,
		Balance:  5000000,
		Currency: model.AccountCurrencyEUR,
	}

	req, err = makeRequest(http.MethodPost, basePath+"/accounts", &createAccountBody)
	require.NoError(t, err)

	var createAccountResp apiv1.AccountResponse

	err = doRequest(req, &createAccountResp)
	require.NoError(t, err)

	require.NotNil(t, createAccountResp)
	assert.Equal(t, createAccountBody.UserID, createAccountResp.UserID)
	assert.Equal(t, createAccountBody.Balance, createAccountResp.Balance)
	assert.Equal(t, createAccountBody.Currency, createAccountResp.Currency)
	assert.Equal(t, createAccountResp.Status, model.AccountStatusActive)
	assert.NotEmpty(t, createAccountResp.ID)
	assert.NotEmpty(t, createAccountResp.Status)
	assert.NotEmpty(t, createAccountResp.Number)

	// Account GET test
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/accounts/", createAccountResp.ID), nil)
	require.NoError(t, err)

	var getAccountResp apiv1.AccountResponse

	err = doRequest(req, &getAccountResp)
	require.NoError(t, err)

	require.NotNil(t, getAccountResp)
	assert.Equal(t, createAccountResp.Currency, getAccountResp.Currency)
	assert.Equal(t, createAccountResp.UserID, getAccountResp.UserID)
	assert.Equal(t, createAccountResp.Status, getAccountResp.Status)
	assert.NotEmpty(t, getAccountResp.ID)
	assert.NotEmpty(t, getAccountResp.UserID)
	assert.NotEmpty(t, getAccountResp.Number)
	assert.NotEmpty(t, getAccountResp.Balance)
	assert.NotEmpty(t, getAccountResp.Currency)
	assert.NotEmpty(t, getAccountResp.Status)

	// Account PUT test
	updateAccountBody := apiv1.AccountUpdateRequest{
		ID:     createAccountResp.ID,
		Status: model.AccountStatusBlocked,
	}
	req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d", basePath, "/accounts/", createAccountResp.ID), &updateAccountBody)
	require.NoError(t, err)

	var updateAccountResp apiv1.AccountResponse

	err = doRequest(req, &updateAccountResp)
	require.NoError(t, err)

	// get after block
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/accounts/", createAccountResp.ID), nil)
	require.NoError(t, err)

	var getStatusAccountResp apiv1.AccountResponse

	err = doRequest(req, &getStatusAccountResp)
	require.NoError(t, err)

	assert.NotNil(t, updateAccountResp)
	assert.Equal(t, getStatusAccountResp.Status, model.AccountStatusBlocked)

	// Delete test user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, &apiv1.UserResponse{})
	require.Error(t, err)
}

func TestCard(t *testing.T) {

	// User TEST POST
	createUserBody := apiv1.UserRequest{
		FullName: "Ugo Chaves",
		Email:    "ven@gmail.com",
		Password: "qwerty12345",
	}

	req, err := makeRequest(http.MethodPost, basePath+"/users", &createUserBody)
	require.NoError(t, err)

	var createUserResp apiv1.UserResponse

	err = doRequest(req, &createUserResp)
	require.NoError(t, err)

	// Account TEST POST
	createAccountBody := apiv1.AccountRequest{
		UserID:   createUserResp.ID,
		Balance:  5000000,
		Currency: model.AccountCurrencyEUR,
	}

	req, err = makeRequest(http.MethodPost, basePath+"/accounts", &createAccountBody)
	require.NoError(t, err)

	var createAccountResp apiv1.AccountResponse

	err = doRequest(req, &createAccountResp)
	require.NoError(t, err)

	// Card POST test
	createCardBody := apiv1.CardRequest{
		AccountID: createAccountResp.ID,
	}
	req, err = makeRequest(http.MethodPost, basePath+"/cards", &createCardBody)
	require.NoError(t, err)

	var createCardResp apiv1.CardResponse

	err = doRequest(req, &createCardResp)
	require.NoError(t, err)

	assert.NotNil(t, createCardResp)
	assert.NotEmpty(t, createCardResp.ID)
	assert.NotEmpty(t, createCardResp.AccountID)
	assert.NotEmpty(t, createCardResp.Number)

	// Card GET test
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/card/", createCardResp.ID), nil)
	require.NoError(t, err)

	var getCardResp apiv1.CardResponse

	err = doRequest(req, &getCardResp)
	require.NoError(t, err)

	assert.NotNil(t, getCardResp)
	assert.NotEmpty(t, getCardResp.ID)
	assert.NotEmpty(t, getCardResp.AccountID)
	assert.NotEmpty(t, getCardResp.Number)
	assert.Equal(t, createAccountResp.ID, getCardResp.AccountID)

	// Delete test user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, &apiv1.UserResponse{})
	require.Error(t, err)
}

func TestPayments(t *testing.T) {

	// User TEST POST
	createUserBody := apiv1.UserRequest{
		FullName: "Fidel Castro",
		Email:    "cub@gmail.com",
		Password: "qwerty12345",
	}

	req, err := makeRequest(http.MethodPost, basePath+"/users", &createUserBody)
	require.NoError(t, err)

	var createUserResp apiv1.UserResponse

	err = doRequest(req, &createUserResp)
	require.NoError(t, err)

	// Account TEST POST
	createAccountBody := apiv1.AccountRequest{
		UserID:   createUserResp.ID,
		Balance:  5000000,
		Currency: model.AccountCurrencyEUR,
	}

	req, err = makeRequest(http.MethodPost, basePath+"/accounts", &createAccountBody)
	require.NoError(t, err)

	var createAccountResp apiv1.AccountResponse

	err = doRequest(req, &createAccountResp)
	require.NoError(t, err)

	// Payments POST test
	createPaymentsBody := apiv1.PaymentRequest{
		AccountID: createAccountResp.ID,
		Sum:       10000,
		Status:    model.PaymentStatusSent,
	}

	createPayAccountBody := apiv1.AccountRequest{
		Balance: createAccountResp.Balance,
	}
	req, err = makeRequest(http.MethodPost, fmt.Sprintf("%s%s%d%s", basePath, "/account/", createAccountResp.ID, "/payment"), createPaymentsBody)
	require.NoError(t, err)

	var createPaymentResp apiv1.PaymentResponse

	err = doRequest(req, &createPaymentResp)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/accounts/", createAccountResp.ID), createPayAccountBody)
	require.NoError(t, err)

	var getAccountResp apiv1.AccountResponse

	err = doRequest(req, &getAccountResp)
	require.NoError(t, err)

	assert.NotNil(t, createPaymentResp)
	assert.NotEmpty(t, createPaymentResp.ID)
	assert.NotEmpty(t, createPaymentResp.AccountID)
	assert.NotEmpty(t, createPaymentResp.Date)
	assert.NotEmpty(t, createPaymentResp.Sum)
	assert.NotEmpty(t, createPaymentResp.Status)
	assert.Equal(t, getAccountResp.Balance, createAccountBody.Balance+createPaymentResp.Sum)

	// Payments ListGET test
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d%s", basePath, "/accounts/", createAccountResp.ID, "/payments"), nil)
	require.NoError(t, err)

	var getListAccountResp []apiv1.PaymentResponse

	err = doRequest(req, &getListAccountResp)
	require.NoError(t, err)

	assert.NotNil(t, getListAccountResp)

	// Delete test user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, &apiv1.UserResponse{})
	require.Error(t, err)
}

func TestRequests(t *testing.T) {

	// User TEST POST
	createUserBody := apiv1.UserRequest{
		FullName: "Sharl de Gol",
		Email:    "fran@gmail.com",
		Password: "qwerty12345",
	}

	req, err := makeRequest(http.MethodPost, basePath+"/users", createUserBody)
	require.NoError(t, err)

	var createUserResp apiv1.UserResponse

	err = doRequest(req, &createUserResp)
	require.NoError(t, err)

	// Account TEST POST
	createAccountBody := apiv1.AccountRequest{
		UserID:   createUserResp.ID,
		Balance:  5000000,
		Currency: model.AccountCurrencyEUR,
	}

	req, err = makeRequest(http.MethodPost, basePath+"/accounts", &createAccountBody)
	require.NoError(t, err)

	var createAccountResp apiv1.AccountResponse

	err = doRequest(req, &createAccountResp)
	require.NoError(t, err)

	// Account TEST POST block
	createBlockAccountBody := apiv1.AccountUpdateRequest{
		ID:     createAccountResp.ID,
		Status: model.AccountStatusBlocked,
	}

	req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d", basePath, "/accounts/", createAccountResp.ID), createBlockAccountBody)
	require.NoError(t, err)

	var updateStatusAccountResp apiv1.AccountResponse

	err = doRequest(req, &updateStatusAccountResp)
	require.NoError(t, err)

	// Requests POST test
	createRequestsBody := apiv1.RequestRequest{
		AccountID: createAccountResp.ID,
	}
	req, err = makeRequest(http.MethodPost, basePath+"/request", createRequestsBody)
	require.NoError(t, err)

	var createRequestResp apiv1.RequestResponse

	err = doRequest(req, &createRequestResp)
	require.NoError(t, err)

	assert.NotEmpty(t, createRequestResp.ID)
	assert.NotEmpty(t, createRequestResp.AccountID)
	assert.NotEmpty(t, createRequestResp.Date)
	assert.NotEmpty(t, createRequestResp.Date)
	assert.Equal(t, createRequestResp.Status, model.RequestStatusNew)
	assert.Equal(t, createRequestResp.AccountID, createAccountResp.ID)

	// Requests PUT approved test
	createRequestApproveBody := apiv1.RequestLockRequest{
		ID:        createRequestResp.ID,
		AccountID: createRequestResp.AccountID,
		Status:    model.RequestStatusApproved,
	}
	req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d%s", basePath, "/request/", createRequestResp.ID, "/approve"), createRequestApproveBody)
	require.NoError(t, err)

	var updateApprovedRequestResp apiv1.RequestResponse

	err = doRequest(req, &updateApprovedRequestResp)
	require.NoError(t, err)

	assert.NotNil(t, updateApprovedRequestResp)
	assert.Equal(t, updateApprovedRequestResp.Status, model.RequestStatusApproved)
	assert.Equal(t, createAccountResp.Status, model.AccountStatusActive)

	// Delete test user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)

	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUserResp.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, &apiv1.UserResponse{})
	require.Error(t, err)
}
