// interaction_service.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// InteractionService representa el servicio de interacción entre los bots A y B
type InteractionService struct {
	botA *BotA
	botB *BotB
}

// NewInteractionService crea una nueva instancia del servicio de interacción
func NewInteractionService(botA *BotA, botB *BotB) *InteractionService {
	return &InteractionService{
		botA: botA,
		botB: botB,
	}
}

// HandleUserRequest maneja la solicitud del usuario y devuelve la respuesta de los bots A y B
func (interactionService *InteractionService) HandleUserRequest(userID, message string) (string, error) {
	// Procesar la solicitud del usuario con el Bot A
	interactionService.botA.ProcessUserResponse(userID, message)

	// Obtener la respuesta del Bot A
	responseA, err := interactionService.botA.GetBotResponse(userID)
	if err != nil {
		return "", err
	}

	// Procesar la respuesta del Bot A con el Bot B
	interactionService.botB.ProcessBotAResponse(userID, responseA)

	// Obtener la respuesta final del Bot B
	responseB, err := interactionService.botB.GetBotResponse(userID)
	if err != nil {
		return "", err
	}

	// Aquí puedes agregar más lógica para manejar la interacción entre los bots A y B

	return responseB, nil
}

// APIHandler maneja las solicitudes de la API del servicio de interacción
func (interactionService *InteractionService) APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		UserID  string `json:"user_id"`
		Message string `json:"message"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.UserID == "" || requestBody.Message == "" {
		http.Error(w, "Invalid user ID or message", http.StatusBadRequest)
		return
	}

	response, err := interactionService.HandleUserRequest(requestBody.UserID, requestBody.Message)
	if err != nil {
		http.Error(w, "Error processing user request", http.StatusInternalServerError)
		return
	}

	responseData := struct {
		Response string `json:"response"`
	}{
		Response: response,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
