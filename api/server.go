package api

import (
	"fmt"
	"log"
	"strings"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/workers"
	"github.com/gorilla/websocket"

	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, clientID string, connections types.ClientConnections) {
	connections[clientID] = types.ClientConnection{SocketConnection: conn}
	worker := workers.NewWorker(types.ClientConnection{SocketConnection: conn})
	worker.Run()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request, connections types.ClientConnections) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	clientID := strings.Split(r.URL.String(), "/")[2]
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws, clientID, connections)
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
