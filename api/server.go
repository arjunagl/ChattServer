package api

import (
	"fmt"
	"log"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, connections types.ClientConnections) {
	connections[uuid.New()] = types.ClientConnection{Connection: conn}
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		// send the message to the client
		msg := []byte("Let's start to talk something.")
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Successfully sent the message to the client")

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request, connections types.ClientConnections) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// helpful log statement to show connections
	log.Println("Client Connected")

	reader(ws, connections)
}

// StartServer Starts the server
func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	connections := make(types.ClientConnections)
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		wsEndpoint(w, r, connections)
	})

	fmt.Println("Listening for incoming connections")
	http.ListenAndServe(":9990", nil)
}
