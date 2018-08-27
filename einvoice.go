package sbank

import "time"

type EInvoiceStatus = string

const (
	EInvoiceStatusAll       EInvoiceStatus = "ALL"
	EInvoiceStatusProcessed                = "PROCESSED"
	EInvoiceStatusDeleted                  = "DELETED"
	EInvoiceStatusNew                      = "NEW"
)

type EInvoicesResponse struct {
	APIResponse
	Items []EInvoice `json:"items"`
}

type EInvoice struct {
	ID                  string    `json:"eFakturaId"`
	IssuerID            string    `json:"issuerId"`
	Reference           string    `json:"eFakturaReference"`
	DocumentType        string    `json:"documentType"`
	Status              string    `json:"status"`
	KID                 string    `json:"kid"`
	OriginalDueDate     time.Time `json:"originalDueDate"`
	OriginalAmount      float64   `json:"originalAmount"`
	MinimumAmount       float64   `json:"minimumAmount"`
	UpdatedDueDate      time.Time `json:"updatedDueDate"`
	UpdatedAmount       float64   `json:"updatedAmount"`
	NotificationDate    time.Time `json:"notificationDate"`
	CreditAccountNumber string    `json:"creditAccountNumber"`
	IssuerName          string    `json:"issuerName"`
}
