package sbank

import (
	"fmt"
	"strings"
)

type CustomersReponse struct {
	Customer     Customer `json:"item"`
	ErrorType    string   `json:"errorType"`
	IsError      bool     `json:"isError"`
	ErrorMessage string   `json:"errorMessage"`
	TraceID      string   `json:"traceId"`
}

// Customer contains all information about a banking customer in sbank.
type Customer struct {
	CustomerID    string        `json:"customerId"`
	FirstName     string        `json:"firstName"`
	LastName      string        `json:"lastName"`
	EmailAddress  string        `json:"emailAddress"`
	DateOfBirth   string        `json:"dateOfBirth"`
	PostalAddress Address       `json:"postalAddress"`
	StreetAddress Address       `json:"streetAddress"`
	PhoneNumbers  []PhoneNumber `json:"phoneNumbers"`
}

func (c Customer) String() string {
	var res []string
	res = append(res, fmt.Sprintf("CustomerID: %s, %s", c.CustomerID))
	res = append(res, fmt.Sprintf("Name: %s, %s", c.FirstName, c.LastName))
	res = append(res, fmt.Sprintf("Email: %s", c.EmailAddress))
	res = append(res, fmt.Sprintf("Birthday: %s", c.DateOfBirth))
	res = append(res, fmt.Sprintf("Postal Address: %s", c.PostalAddress))
	res = append(res, fmt.Sprintf("Street Address: %s", c.PostalAddress))

	switch len(c.PhoneNumbers) {
	case 0:
		// Do nothing
	case 1:
		res = append(res, fmt.Sprintf("Phone number: %s", c.PhoneNumbers[0]))
	default:
		res = append(res, "Phone numbers:")
		for _, pn := range c.PhoneNumbers {
			res = append(res, pn.String())
		}
	}

	return strings.Join(res, "\n")
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

func (p PhoneNumber) String() string {
	var res string
	if s := p.CountryCode; s != "" {
		res = fmt.Sprintf("(+%s) ", s)
	}
	res += fmt.Sprintf("%s", p.Number)
	return res
}

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	AddressLine3 string `json:"addressLine3"`
	AddressLine4 string `json:"addressLine4"`
	Country      string `json:"country"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
}

func (a Address) String() string {
	var res []string

	if s := a.AddressLine1; s != "" {
		res = append(res, s)
	}
	if s := a.AddressLine2; s != "" {
		res = append(res, s)
	}
	if s := a.AddressLine3; s != "" {
		res = append(res, s)
	}
	if s := a.AddressLine4; s != "" {
		res = append(res, s)
	}
	if s := a.ZipCode; s != "" {
		res = append(res, s)
	}
	if s := a.City; s != "" {
		res = append(res, s)
	}
	if s := a.Country; s != "" {
		res = append(res, s)
	}
	return strings.Join(res, ", ")
}
