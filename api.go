package sbank

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	apiURL = "https://api.sbanken.no"
)

// Client can be used to communicate with the sbank APIs.
// https://sbanken.no/bruke/utviklerportalen/
type Client struct {
	client *http.Client

	userID string
}

// NewWithClient returns a new Client using the provided http.Client as
// underlying transport. Note that the provided http.Client must handle the
// authorization.
func NewWithClient(uid string, c *http.Client) *Client {
	return &Client{
		client: c,
		userID: uid,
	}
}

// New returns a new Client using the provided uid, id and secret for
// authentication.
func New(uid, id, secret string) *Client {
	ccreds := clientcredentials.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     "https://api.sbanken.no/identityserver/connect/token",
	}

	return &Client{
		client: ccreds.Client(context.Background()),
		userID: uid,
	}
}

// Accounts lists the accounts owned by the customer and the accounts the
// customer has been granted access to.
func (a *Client) Accounts() (AccountsResponse, error) {
	var res AccountsResponse

	path := fmt.Sprintf("%s/Bank/api/v1/Accounts", apiURL)
	if err := get(a, path, &res); err != nil {
		return res, fmt.Errorf("accounts: %s", err)
	}

	return res, nil
}

// Account reads an account owned by the customer or an account that the
// customer has been granted access to.
func (a *Client) Account(uid, aid string) (AccountResponse, error) {
	var res AccountResponse

	path := fmt.Sprintf("%s/Bank/api/v1/Accounts/%s/%s", apiURL, uid, aid)
	if err := get(a, path, &res); err != nil {
		return res, fmt.Errorf("account: %s", err)
	}

	return res, nil
}

// EInvoices is the default listing method for current not-processed/new
// invoices.
//
// Note: We can only show a maximal of 16 months back in time.
func (a *Client) EInvoices(status EInvoiceStatus, from, to *time.Time, offset, limit *int) (EInvoicesResponse, error) {
	var res EInvoicesResponse

	v := url.Values{}
	v.Add("status", status)
	if from != nil {
		v.Add("from", from.Format(time.RFC3339))
	}
	if to != nil {
		v.Add("to", to.Format(time.RFC3339))
	}
	if offset != nil {
		v.Add("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		v.Add("limit", strconv.Itoa(*limit))
	}

	path := fmt.Sprintf("%s/Bank/api/v1/Efakturas?%s", apiURL, v.Encode())
	if err := get(a, path, &res); err != nil {
		return res, fmt.Errorf("einvoices: %s", err)
	}

	return res, nil
}

// Transactions returns the latest transactions of the given account within the
// time span set by the start and end date parameters.
//
// Note that dateTime type parameters are relative to Central European Time
// (GMT+1); only the date part is relevant.
//
// * from/to: the range of dates to retrieve transactions for.
// * offset/limit: pagination for results.
func (a *Client) Transactions(aid string, from, to *time.Time, offset, limit *int) (TransactionsResponse, error) {
	var res TransactionsResponse

	v := url.Values{}
	if from != nil {
		v.Add("from", from.Format(time.RFC3339))
	}
	if to != nil {
		v.Add("to", to.Format(time.RFC3339))
	}
	if offset != nil {
		v.Add("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		v.Add("limit", strconv.Itoa(*limit))
	}

	path := fmt.Sprintf("%s/Bank/api/v1/Transactions/%s?%s", apiURL, aid, v.Encode())
	if err := get(a, path, &res); err != nil {
		return res, fmt.Errorf("transactions: %s", err)
	}

	return res, nil
}

func get(c *Client, url string, res interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to build request: %s", url, err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("CustomerID", c.userID)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed to perform request: %s", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: expected status %d got %d", url, http.StatusOK, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return fmt.Errorf("%s: failed to decode request: %s", url, err)
	}

	return nil
}
