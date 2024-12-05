// server.go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Definição do Upgrader do WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permite conexões de qualquer origem
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Atualiza a conexão para WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		currentTimeFormatted := time.Now().Format("02-01-2006 15:04:05")
		clientId := r.URL.Query().Get("userID")

		// Envia uma mensagem a cada 60 segundos
		message := "A atualização foi realizada para o usuário " + clientId + " : " + currentTimeFormatted
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(currentTimeFormatted + " - send message")

		// Espera 60 segundos antes de enviar outra mensagem
		// Simulando uma atualização a cada 60 segundos
		<-time.After(60 * time.Second)
	}
}

func main() {
	// Rota WebSocket
	http.HandleFunc("/ws", handler)

	// Inicia o servidor na porta 8080
	log.Println("Servidor WebSocket iniciado na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
