package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/arjunagl/ChattServer/api/workers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var workerChannels = make(types.WorkerChannels)

func startWorker(conn *websocket.Conn, clientID string) {
	worker := workers.NewWorker(clientID, types.ClientConnection{SocketConnection: conn}, workerChannels)
	worker.Run()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	clientID := mux.Vars(r)["clientID"]

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	startWorker(ws, clientID)
}

func handleSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling subscription")
	clientID := mux.Vars(r)["clientID"]

	s := &webpush.Subscription{}
	if decodeErr := json.NewDecoder(r.Body).Decode(s); decodeErr != nil && decodeErr != io.EOF {
		fmt.Println("Error decoding subscription")
	}
	workerChannels[clientID] <- commands.WorkerCommand{Command: commands.SetClientSubscription, Details: commands.SetClientSubscriptionCommandDetails{ClientSubscription: s}}
	w.WriteHeader(200)
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

	r.HandleFunc("/subscribe/{clientID}", func(w http.ResponseWriter, r *http.Request) {
		handleSubscription(w, r)
	}).Methods("POST", "OPTIONS")

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"https://localhost:3001"}),
		handlers.AllowCredentials(),
	)
	r.Use(cors)

	http.ListenAndServe(":9990", r)
}
