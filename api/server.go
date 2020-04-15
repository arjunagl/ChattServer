package api

import (
	"fmt"
	"log"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/workers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var workerChannels = make(types.WorkerChannels)

func reader(conn *websocket.Conn, clientID string) {
	worker := workers.NewWorker(types.ClientConnection{SocketConnection: conn, CientID: clientID}, workerChannels)
	worker.Run()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Web Socket endpoint hit")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	clientID := mux.Vars(r)["clientID"]

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws, clientID)
}

// StartServer Starts the server
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	r.HandleFunc("/ws/{clientID}", func(w http.ResponseWriter, r *http.Request) {
		wsEndpoint(w, r)
	})

	r.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(":9990", nil)
}
