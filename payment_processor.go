// payment_processor.go

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PaymentProcessor representa el módulo que procesa los pagos con Bitcoin y Lightning Network
type PaymentProcessor struct {
	lnbitsAPIKey string // API key para la integración con LNbits
}

// NewPaymentProcessor crea una nueva instancia del PaymentProcessor
func NewPaymentProcessor(lnbitsAPIKey string) *PaymentProcessor {
	return &PaymentProcessor{
		lnbitsAPIKey: lnbitsAPIKey,
	}
}

// ProcessPayment procesa un pago con Bitcoin y Lightning Network utilizando LNbits API
func (paymentProcessor *PaymentProcessor) ProcessPayment(userID string, amount float64) (string, error) {
	// Verificar que el monto del pago sea válido
	if amount <= 0 {
		return "", errors.New("invalid payment amount")
	}

	// Realizar el pago utilizando LNbits API
	paymentID, err := paymentProcessor.makePayment(userID, amount)
	if err != nil {
		return "", fmt.Errorf("payment failed: %v", err)
	}

	return paymentID, nil
}

// makePayment realiza el pago utilizando la API de LNbits
func (paymentProcessor *PaymentProcessor) makePayment(userID string, amount float64) (string, error) {
	// Realizar la solicitud HTTP a la API de LNbits para generar un pago
	url := fmt.Sprintf("https://lnbits.com/api/v1/payments")
	data := fmt.Sprintf(`{"out": true, "amount": %f, "memo": "Payment for user: %s"}`, amount, userID)
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	request.Header.Set("X-Api-Key", paymentProcessor.lnbitsAPIKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("payment request failed")
	}

	// Leer el ID de pago de la respuesta de la API
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var paymentResponse map[string]interface{}
	err = json.Unmarshal(body, &paymentResponse)
	if err != nil {
		return "", err
	}

	paymentID, ok := paymentResponse["payment_hash"].(string)
	if !ok {
		return "", errors.New("invalid payment response")
	}

	return paymentID, nil
}
