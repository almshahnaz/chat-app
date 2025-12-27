package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading: ", err)
		return
	}
	defer conn.Close()

	// Listen for incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break	
		}
		fmt.Printf("Received: %s\n", message)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message: ", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}