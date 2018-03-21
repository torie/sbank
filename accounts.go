package sbank

import (
	"fmt"
	"strings"
)

type AccountsResponse struct {
	AvailableItems int           `json:"availableItems"`
	Accounts       []AccountInfo `json:"items"`
	ErrorType      string        `json:"errorType"`
	IsError        bool          `json:"isError"`
	ErrorMessage   string        `json:"errorMessage"`
	TraceID        string        `json:"traceId"`
}

type AccountResponse struct {
	Account      AccountInfo `json:"item"`
	ErrorType    string      `json:"errorType"`
	IsError      bool        `json:"isError"`
	ErrorMessage string      `json:"errorMessage"`
	TraceID      string      `json:"traceId"`
}

type AccountInfo struct {
	AccountNumber   string  `json:"accountNumber"`
	CustomerID      string  `json:"customerId"`
	OwnerCustomerID string  `json:"ownerCustomerId"`
	Name            string  `json:"name"`
	AccountType     string  `json:"accountType"`
	Available       float64 `json:"available"`
	Balance         float64 `json:"balance"`
	CreditLimit     float64 `json:"creditLimit"`
	DefaultAccount  bool    `json:"defaultAccount"`
}

func (a AccountInfo) String() string {
	var res []string
	res = append(res, fmt.Sprintf("Account number: %s", a.AccountNumber))
	res = append(res, fmt.Sprintf("Customer ID: %s", a.CustomerID))
	res = append(res, fmt.Sprintf("Owner customer ID: %s", a.OwnerCustomerID))
	res = append(res, fmt.Sprintf("Name: %s", a.Name))
	res = append(res, fmt.Sprintf("Account type: %s", a.AccountType))
	res = append(res, fmt.Sprintf("Available: %f", a.Available))
	res = append(res, fmt.Sprintf("Balance: %f", a.Balance))
	res = append(res, fmt.Sprintf("Credit limit: %f", a.CreditLimit))

	dac := "no"
	if a.DefaultAccount {
		dac = "yes"
	}
	res = append(res, fmt.Sprintf("Default account: %s", dac))
	return strings.Join(res, "\n")
}
