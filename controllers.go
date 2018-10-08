// Implements a jsonapi for http responses
package jsonapi

import (
	"encoding/json"
	"github.com/quorumsco/logs"
	"net/http"
)

// Returns a Success response
func Success(w http.ResponseWriter, req *http.Request, data interface{}, status int) {
	if data == nil {
		w.WriteHeader(status)
		return
	}

	b, err := json.Marshal(SuccessView{Status: "success", Data: data})
	if err != nil {
		logs.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Returns a Success response if the data is not nil
func SuccessOKOr404(w http.ResponseWriter, req *http.Request, data interface{}) {
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(SuccessView{Status: "success", Data: data})
	if err != nil {
		logs.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// Returns a Fail response
func Fail(w http.ResponseWriter, req *http.Request, data interface{}, status int) {
	b, err := json.Marshal(FailView{Status: "fail", Data: data})
	if err != nil {
		logs.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Returns an Error response
func Error(w http.ResponseWriter, req *http.Request, message string, status int) {
	b, err := json.Marshal(ErrorView{Status: "error", Message: message})
	if err != nil {
		logs.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Unmarshaled json data into the passed interface
func Request(data interface{}, req *http.Request) error {
	var r = new(RequestView)
	var err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(r.Data, data)
	if err != nil {
		return err
	}
	return nil
}
