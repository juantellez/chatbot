// value_provider.go

package main

import (
	"errors"
)

// ValueProvider representa el módulo que proporciona valor al usuario
type ValueProvider struct {
	botA *BotA
	botB *BotB
}

// NewValueProvider crea una nueva instancia del ValueProvider
func NewValueProvider(botA *BotA, botB *BotB) *ValueProvider {
	return &ValueProvider{
		botA: botA,
		botB: botB,
	}
}

// ProvideValue proporciona valor al usuario basado en la interacción con los bots A y B
func (valueProvider *ValueProvider) ProvideValue(userID, message string) (string, error) {
	// Procesar el mensaje con el bot A
	valueProvider.botA.ProcessUserResponse(userID, message)

	// Obtener la respuesta del bot A
	responseA, err := valueProvider.botA.GetBotResponse(userID)
	if err != nil {
		return "", err
	}

	// Procesar la respuesta del bot A con el bot B
	valueProvider.botB.ProcessBotAResponse(userID, responseA)

	// Obtener la respuesta final del bot B
	responseB, err := valueProvider.botB.GetBotResponse(userID)
	if err != nil {
		return "", err
	}

	// Verificar si se ha completado el proceso de KYC del usuario
	user, err := valueProvider.botA.findOrCreateUser(userID)
	if err != nil {
		return "", err
	}

	if !user.IsKYCDone {
		return "", errors.New("KYC process not completed")
	}

	return responseB, nil
}
