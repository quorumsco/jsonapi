package jsonapi

import "encoding/json"

type RequestView struct {
	Data json.RawMessage `json:"data"`
}

type SuccessView struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type FailView struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorView struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
