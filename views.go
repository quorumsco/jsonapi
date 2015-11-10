// Implements a jsonapi for http responses
package jsonapi

import "encoding/json"

// RequestView represents the json data received
type RequestView struct {
	Data json.RawMessage `json:"data"`
}

// SuccessView represent a success json response
type SuccessView struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// SuccessView represent a fail json response
type FailView struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// SuccessView represent an error json response
type ErrorView struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
