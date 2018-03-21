package sbank

import (
	"fmt"
	"strings"
	"time"
)

type TransactionsResponse struct {
	AvailableItems int           `json:"availableItems"`
	Transactions   []Transaction `json:"items"`
	ErrorType      string        `json:"errorType"`
	IsError        bool          `json:"isError"`
	ErrorMessage   string        `json:"errorMessage"`
	TraceID        string        `json:"traceId"`
}

type Transaction struct {
	TransactionID    string    `json:"transactionId"`
	CustomerID       string    `json:"customerId"`
	AccountNumber    string    `json:"accountNumber"`
	Amount           float64   `json:"amount"`
	Text             string    `json:"text"`
	TransactionType  string    `json:"transactionType"`
	RegistrationDate time.Time `json:"registrationDate"`
	AccountingDate   time.Time `json:"accountingDate"`
	InterestDate     time.Time `json:"interestDate"`
}

func (t Transaction) String() string {
	var res []string
	res = append(res, fmt.Sprintf("Transaction id: %s", t.TransactionID))
	res = append(res, fmt.Sprintf("Customer id: %s", t.CustomerID))
	res = append(res, fmt.Sprintf("Account numer: %s", t.AccountNumber))
	res = append(res, fmt.Sprintf("Amount: %f", t.Amount))
	res = append(res, fmt.Sprintf("Text: %s", t.Text))
	res = append(res, fmt.Sprintf("Transaction type: %s", t.TransactionType))

	res = append(res, fmt.Sprintf("Registration date: %s", t.RegistrationDate.Format("2006-01-02")))
	res = append(res, fmt.Sprintf("Accounting date: %s", t.AccountingDate.Format("2006-01-02")))
	res = append(res, fmt.Sprintf("Interest date: %s", t.InterestDate.Format("2006-01-02")))

	return strings.Join(res, "\n")
}
