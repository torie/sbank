package sbank

type APIResponse struct {
	AvailableItems int    `json:"availableItems"`
	ErrorType      string `json:"errorType"`
	IsError        bool   `json:"isError"`
	ErrorMessage   string `json:"errorMessage"`
	TraceID        string `json:"traceId"`
}
