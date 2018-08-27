package sbank

import (
	"fmt"
	"strings"
	"time"
)

type TransactionsResponse struct {
	APIResponse
	Items []Transaction `json:"items"`
}

type Transaction struct {
	AccountingDate              time.Time   `json:"accountingDate"`
	InterestDate                time.Time   `json:"interestDate"`
	OtherAccountNumber          string      `json:"otherAccountNumber"`
	OtherAccountNumberSpecified bool        `json:"otherAccountNumberSpecified"`
	Amount                      float64     `json:"amount"`
	Text                        string      `json:"text"`
	TransactionType             string      `json:"transactionType"`
	TransactionTypeCode         int         `json:"transactionTypeCode"`
	TransactionTypeText         string      `json:"transactionTypeText"`
	IsReservation               bool        `json:"isReservation"`
	ReservationType             string      `json:"reservationType"`
	Source                      int         `json:"source"`
	CardDetailsSpecified        bool        `json:"cardDetailsSpecified"`
	CardDetails                 CardDetails `json:"cardDetails"`
}

func (t Transaction) String() string {
	var res []string

	res = append(res, fmt.Sprintf("AccountingDate: %s", t.AccountingDate.Format("2006-01-02")))
	res = append(res, fmt.Sprintf("InterestDate: %s", t.InterestDate.Format("2006-01-02")))
	res = append(res, fmt.Sprintf("OtherAccountNumber: %s", t.OtherAccountNumber))
	res = append(res, fmt.Sprintf("OtherAccountNumberSpecified: %t", t.OtherAccountNumberSpecified))
	res = append(res, fmt.Sprintf("Amount: %f", t.Amount))
	res = append(res, fmt.Sprintf("Text: %s", t.Text))
	res = append(res, fmt.Sprintf("TransactionType: %s", t.TransactionType))
	res = append(res, fmt.Sprintf("TransactionTypeCode: %d", t.TransactionTypeCode))
	res = append(res, fmt.Sprintf("TransactionTypeText: %s", t.TransactionTypeText))
	res = append(res, fmt.Sprintf("IsReservation: %t", t.IsReservation))
	res = append(res, fmt.Sprintf("ReservationType: %s", t.ReservationType))
	res = append(res, fmt.Sprintf("Source: %d", t.Source))
	if t.CardDetailsSpecified {
		res = append(res, fmt.Sprintf("  CardNumber: %s", t.CardDetails.CardNumber))
		res = append(res, fmt.Sprintf("  CurrencyAmount: %f", t.CardDetails.CurrencyAmount))
		res = append(res, fmt.Sprintf("  CurrencyRate: %f", t.CardDetails.CurrencyRate))
		res = append(res, fmt.Sprintf("  MerchantCategoryCode: %s", t.CardDetails.MerchantCategoryCode))
		res = append(res, fmt.Sprintf("  MerchantCategoryDescription: %s", t.CardDetails.MerchantCategoryDescription))
		res = append(res, fmt.Sprintf("  MerchantCity: %s", t.CardDetails.MerchantCity))
		res = append(res, fmt.Sprintf("  MerchantName: %s", t.CardDetails.MerchantName))
		res = append(res, fmt.Sprintf("  OriginalCurrencyCode: %s", t.CardDetails.OriginalCurrencyCode))
		res = append(res, fmt.Sprintf("  PurchaseDate: %s", t.CardDetails.PurchaseDate.Format("2006-01-02")))
		res = append(res, fmt.Sprintf("  TransactionID: %s", t.CardDetails.TransactionID))
	}

	return strings.Join(res, "\n")
}

type CardDetails struct {
	CardNumber                  string    `json:"cardNumber"`
	CurrencyAmount              float64   `json:"currencyAmount"`
	CurrencyRate                float64   `json:"currencyRate"`
	MerchantCategoryCode        string    `json:"merchantCategoryCode"`
	MerchantCategoryDescription string    `json:"merchantCategoryDescription"`
	MerchantCity                string    `json:"merchantCity"`
	MerchantName                string    `json:"merchantName"`
	OriginalCurrencyCode        string    `json:"originalCurrencyCode"`
	PurchaseDate                time.Time `json:"purchaseDate"`
	TransactionID               string    `json:"transactionId"`
}
