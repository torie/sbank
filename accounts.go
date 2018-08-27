package sbank

import (
	"fmt"
	"strings"
)

type AccountsResponse struct {
	APIResponse
	Items []AccountInfo `json:"items"`
}

type AccountResponse struct {
	Account      AccountInfo `json:"item"`
	ErrorType    string      `json:"errorType"`
	IsError      bool        `json:"isError"`
	ErrorMessage string      `json:"errorMessage"`
	TraceID      string      `json:"traceId"`
}

type AccountInfo struct {
	AccountID       string  `json:"accountID"`
	AccountNumber   string  `json:"accountNumber"`
	OwnerCustomerID string  `json:"ownerCustomerId"`
	Name            string  `json:"name"`
	AccountType     string  `json:"accountType"`
	Available       float64 `json:"available"`
	Balance         float64 `json:"balance"`
	CreditLimit     float64 `json:"creditLimit"`
}

func (a AccountInfo) String() string {
	var res []string
	res = append(res, fmt.Sprintf("Account id: %s", a.AccountID))
	res = append(res, fmt.Sprintf("Account number: %s", a.AccountNumber))
	res = append(res, fmt.Sprintf("Owner customer ID: %s", a.OwnerCustomerID))
	res = append(res, fmt.Sprintf("Name: %s", a.Name))
	res = append(res, fmt.Sprintf("Account type: %s", a.AccountType))
	res = append(res, fmt.Sprintf("Available: %f", a.Available))
	res = append(res, fmt.Sprintf("Balance: %f", a.Balance))
	res = append(res, fmt.Sprintf("Credit limit: %f", a.CreditLimit))

	return strings.Join(res, "\n")
}
