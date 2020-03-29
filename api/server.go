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

var workerChannels = make(types.WorkerChannels)

func reader(conn *websocket.Conn, clientID string) {
	worker := workers.NewWorker(types.ClientConnection{SocketConnection: conn, CientID: clientID})

	worker.Run()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	clientID := strings.Split(r.URL.String(), "/")[2]
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws, clientID)
}

// StartServer Starts the server
func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	connections := make(types.ClientConnections)
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		wsEndpoint(w, r)
	})

	http.ListenAndServe(":9990", nil)
}
