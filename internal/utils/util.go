package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Error struct {
	StatusCode    int    `json:"statusCode"`
	Error         string `json:"error"`
	MessageToUser string `json:"messageToUser"`
}

func ParseBody(body io.Reader, out interface{}) error {
	err := json.NewDecoder(body).Decode(out)
	if err != nil {
		return err
	}

	return nil
}

func EncodeJSONBody(resp http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(resp).Encode(data)
}

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json") // accept for post use
	w.WriteHeader(statusCode)

	if body != nil {
		if err := EncodeJSONBody(w, body); err != nil {
			fmt.Printf("failed to respond JSON with error: %v", err)
		}
	}
}

func RespondError(w http.ResponseWriter, statusCode int, err error, messageToUser string) {
	w.WriteHeader(statusCode)

	var errString string
	if err != nil {
		errString = err.Error()
	}

	newError := Error{
		MessageToUser: messageToUser,
		Error:         errString,
		StatusCode:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(newError); err != nil {
		fmt.Printf("failed to send error to caller with error: %v", err)
	}
}
