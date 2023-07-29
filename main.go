// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Cargar la configuración del sistema desde el archivo "config.json"
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Establecer la conexión a la base de datos
	database, err := NewDatabase(config.DatabasePath)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Close()

	// Crear las instancias de los bots A y B
	botA := NewBotA()
	botB := NewBotB()

	// Crear el servicio de interacción entre los bots A y B
	interactionService := NewInteractionService(botA, botB)

	// Crear el módulo ValueProvider para proporcionar valor al usuario
	valueProvider := NewValueProvider(botA, botB)

	// Crear el módulo PaymentProcessor para procesar pagos con Bitcoin y LNbits
	paymentProcessor := NewPaymentProcessor(config.LNbitsAPIKey)

	// Crear el servidor HTTP para manejar las solicitudes del usuario
	http.HandleFunc("/api/interact", interactionService.APIHandler)

	// Ejemplo de manejo de una solicitud del usuario y procesamiento de pago
	http.HandleFunc("/api/payment", func(w http.ResponseWriter, r *http.Request) {
		// Obtener el userID y el monto del pago desde el request
		userID := r.FormValue("user_id")
		amount := r.FormValue("amount")

		// Realizar el pago con el monto proporcionado utilizando el módulo PaymentProcessor
		paymentID, err := paymentProcessor.ProcessPayment(userID, amount)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing payment: %v", err), http.StatusInternalServerError)
			return
		}

		// Responder al cliente con el ID del pago generado
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Payment ID: %s", paymentID)))
	})

	// Iniciar el servidor HTTP
	address := "localhost:8080"
	fmt.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
