// bot_a.go

package main

import (
	"errors"
	"sync"
)

// BotA representa el Bot A
type BotA struct {
	users map[string]*User
	mutex sync.RWMutex
}

// NewBotA crea una nueva instancia del Bot A
func NewBotA() *BotA {
	return &BotA{
		users: make(map[string]*User),
	}
}

// ProcessUserResponse procesa la respuesta del usuario y guarda la información relevante
func (botA *BotA) ProcessUserResponse(userID, message string) {
	botA.mutex.Lock()
	defer botA.mutex.Unlock()

	user, ok := botA.users[userID]
	if !ok {
		user = &User{
			UserID: userID,
		}
		botA.users[userID] = user
	}

	// Procesar la respuesta del usuario y guardar información relevante
	user.LastInteraction = message
	user.IsKYCDone = false // Asignar false temporalmente hasta que se complete el proceso de KYC

	// Aquí puedes agregar más lógica para procesar y almacenar información del usuario
}

// GetBotResponse obtiene la respuesta del Bot A al usuario
func (botA *BotA) GetBotResponse(userID string) (string, error) {
	botA.mutex.RLock()
	defer botA.mutex.RUnlock()

	user, ok := botA.users[userID]
	if !ok {
		return "", errors.New("user not found")
	}

	// Aquí puedes implementar la lógica para generar la respuesta del Bot A al usuario
	response := "Hola, soy el Bot A. ¿En qué puedo ayudarte hoy?"

	return response, nil
}

// findOrCreateUser busca o crea un usuario en el Bot A
func (botA *BotA) findOrCreateUser(userID string) (*User, error) {
	botA.mutex.Lock()
	defer botA.mutex.Unlock()

	user, ok := botA.users[userID]
	if !ok {
		user = &User{
			UserID: userID,
		}
		botA.users[userID] = user
	}

	return user, nil
}
