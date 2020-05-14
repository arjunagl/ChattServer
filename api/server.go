package api

import (
	"fmt"
	"log"

	"net/http"

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

func handleSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling subscription")
	// s := &webpush.Subscription{}
	// json.Unmarshal([]byte("<YOUR_SUBSCRIPTION>"), s)

	// // Send Notification
	// resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
	// 	Subscriber:      "chatt-server@chatt-server.com",
	// 	VAPIDPublicKey:  "BM221uCcUB6tJBektDBpuhrFtvECNs7mcShfG6NUnUUR1lV7vGWmWMm7eNZ0ztW4IjDPsGOAG9sQOkjP1hC_23A",
	// 	VAPIDPrivateKey: "9LhvZAWJpanJGmkhA416muEYCWOyqzCbV_5P-Z_WR-c",
	// 	TTL:             30,
	// })
	// if err != nil {
	// 	fmt.Printf("Error sending push notification = %v", err)
	// }
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.WriteHeader(200)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// defer resp.Body.Close()
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
