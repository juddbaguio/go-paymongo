package paymongo

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

var PAYMONGO_API_URL = "https://api.paymongo.com/v1"

func PaymongoAuthHeader(secret string) string {
	return fmt.Sprintf("Basic %s", secret)
}

func PaymongoApiEndpoint(base string, path string) string {
	return fmt.Sprintf("%s%s", base, path)
}

func (p *PaymongoInstance) makeRequest(method string, path string, payload *bytes.Buffer) (*http.Response, error) {
	var client = &http.Client{}
	var req *http.Request
	var err error

	if payload != nil {
		req, err = http.NewRequest(method, PaymongoApiEndpoint(PAYMONGO_API_URL, path), payload)
	} else {
		req, err = http.NewRequest(method, PaymongoApiEndpoint(PAYMONGO_API_URL, path), nil)
	}

	if err != nil {
		return nil, errors.New("invalid request")
	}

	req.Header.Add("Authorization", PaymongoAuthHeader(p.secret))
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}
