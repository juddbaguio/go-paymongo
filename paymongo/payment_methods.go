package paymongo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CardDetails struct {
	CardNumber string `json:"card_number"`
	ExpMonth   int    `json:"exp_month"`
	ExpYear    int    `json:"exp_year"`
	Cvc        string `json:"cvc"`
}

type CreatePaymentPayloadAttributes struct {
	Type    string      `json:"type"`
	Details CardDetails `json:"details"`
}

type CreatePaymentData struct {
	Attributes CreatePaymentPayloadAttributes `json:"attributes"`
}

type CreatePaymentPayload struct {
	Data CreatePaymentData `json:"data"`
}

func (p *PaymongoInstance) CreatePaymentMethod(data CreatePaymentPayload) (map[string]interface{}, error) {
	jsonData, _ := json.Marshal(data)
	responseData, err := p.makeRequest("POST", "/payment_methods", bytes.NewBuffer(jsonData))

	var paymentMethodResponse map[string]interface{}

	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(responseData.Body)

	json.Unmarshal(buf.Bytes(), &paymentMethodResponse)
	return paymentMethodResponse, nil
}

func (p *PaymongoInstance) RetrievePaymentMethod(paymentMethodId string) (map[string]interface{}, error) {
	paymentMethodIdPath := fmt.Sprintf("/payment_methods/%s", paymentMethodId)
	responseData, err := p.makeRequest("GET", paymentMethodIdPath, nil)

	var paymentMethod map[string]interface{}

	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(responseData.Body)

	json.Unmarshal(buf.Bytes(), &paymentMethod)
	return paymentMethod, nil
}
