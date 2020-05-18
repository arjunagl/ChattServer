package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/arjunagl/ChattServer/api/types"
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

func reader(conn *websocket.Conn, clientID string) {
	worker := workers.NewWorker(clientID, types.ClientConnection{SocketConnection: conn}, workerChannels)
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

func handleSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling subscription")
	clientID := mux.Vars(r)["clientID"]
	goWorker := workers[clientID]
	s := &webpush.Subscription{}
	if decodeErr := json.NewDecoder(r.Body).Decode(s); decodeErr != nil && decodeErr != io.EOF {
		fmt.Println("Error decoding subscription")
	}

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Push subscription successful"), s, &webpush.Options{
		Subscriber:      "chatt-server@chatt-server.com",
		VAPIDPublicKey:  "BM221uCcUB6tJBektDBpuhrFtvECNs7mcShfG6NUnUUR1lV7vGWmWMm7eNZ0ztW4IjDPsGOAG9sQOkjP1hC_23A",
		VAPIDPrivateKey: "9LhvZAWJpanJGmkhA416muEYCWOyqzCbV_5P-Z_WR-c",
		TTL:             30,
	})
	if err != nil {
		fmt.Printf("Error sending push notification = %v", err)
	}
	defer resp.Body.Close()
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
