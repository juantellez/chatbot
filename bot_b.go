// bot_b.go

package main

import (
	"errors"
	"sync"
)

// BotB representa el Bot B
type BotB struct {
	users map[string]*User
	mutex sync.RWMutex
}

// NewBotB crea una nueva instancia del Bot B
func NewBotB() *BotB {
	return &BotB{
		users: make(map[string]*User),
	}
}

// ProcessBotAResponse procesa la respuesta del Bot A y guarda la información relevante
func (botB *BotB) ProcessBotAResponse(userID, responseA string) {
	botB.mutex.Lock()
	defer botB.mutex.Unlock()

	user, ok := botB.users[userID]
	if !ok {
		user = &User{
			UserID: userID,
		}
		botB.users[userID] = user
	}

	// Procesar la respuesta del Bot A y guardar información relevante
	user.LastBotAResponse = responseA

	// Aquí puedes agregar más lógica para procesar y almacenar información relevante del Bot A
}

// GetBotResponse obtiene la respuesta del Bot B al usuario
func (botB *BotB) GetBotResponse(userID string) (string, error) {
	botB.mutex.RLock()
	defer botB.mutex.RUnlock()

	user, ok := botB.users[userID]
	if !ok {
		return "", errors.New("user not found")
	}

	// Aquí puedes implementar la lógica para generar la respuesta del Bot B al usuario
	response := "Hola, soy el Bot B. He recibido esta respuesta del Bot A: " + user.LastBotAResponse

	return response, nil
}
