package sbank

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
}

// NewWithClient returns a new Client using the provided http.Client as
// underlying transport. Note that the provided http.Client must handle the
// authorization.
func NewWithClient(c *http.Client) *Client {
	return &Client{
		client: c,
	}
}

// New returns a new Client using the provided id and secret for authentication.
func New(id, secret string) *Client {
	// FIXME: We are currently using a fork of golang.org/x/oauth2 because the
	// sbank API does not correctly implement RFC6749. Specifically, the id. and
	// secret should be urlencoded before being combined and base64 encoded
	// (https://tools.ietf.org/html/rfc6749#section-2.3.1).
	//
	// However, the sbank API requires the id and secret to not be urlencoded.
	// This is handled by the golang.org/x/oauth2 by adding the token url to a
	// list of "bad providers". Until this is fixed in golang.org/x/oauth2, the
	// current fork handles this.
	ccreds := clientcredentials.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     "https://api.sbanken.no/identityserver/connect/token",
	}

	return &Client{
		client: ccreds.Client(context.Background()),
	}
}

// Accounts returns the account information for all accounts which belongs to
// the provided user.
func (a *Client) Accounts(uid string) (AccountsResponse, error) {
	var res AccountsResponse

	path := fmt.Sprintf("%s/Bank/api/v1/Accounts/%s", apiURL, uid)
	if err := get(a.client, path, &res); err != nil {
		return res, fmt.Errorf("accounts: failed to build request: %s", err)
	}

	return res, nil
}

// Account returns the account information for a single account.
func (a *Client) Account(uid, aid string) (AccountResponse, error) {
	var res AccountResponse

	path := fmt.Sprintf("%s/Bank/api/v1/Accounts/%s/%s", apiURL, uid, aid)
	if err := get(a.client, path, &res); err != nil {
		return res, fmt.Errorf("account: failed to build request: %s", err)
	}

	return res, nil
}

// Transactions returns transactions for a single account. Additional parameters
// can be provided to limit the results returned.
//
// * from/to: the range of dates to retrieve transactions for.
// * offset/limit: pagination for results.
func (a *Client) Transactions(uid, aid string, from, to *time.Time, offset, limit *int) (TransactionsResponse, error) {
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

	path := fmt.Sprintf("%s/Bank/api/v1/Transactions/%s/%s%s", apiURL, uid, aid, v.Encode())
	log.Println(path)
	if err := get(a.client, path, &res); err != nil {
		return res, fmt.Errorf("transactions: failed to build request: %s", err)
	}

	return res, nil
}

// Customer returns information about a bank customer.
func (a *Client) Customer(uid string) (CustomersReponse, error) {
	var res CustomersReponse

	path := fmt.Sprintf("%s/Customers/api/v1/Customers/%s", apiURL, uid)
	if err := get(a.client, path, &res); err != nil {
		return res, fmt.Errorf("customers: failed to build request: %s", err)
	}

	return res, nil
}

func get(c *http.Client, url string, res interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to build request: %s", url, err)
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed to peform request: %s", url, err)
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
