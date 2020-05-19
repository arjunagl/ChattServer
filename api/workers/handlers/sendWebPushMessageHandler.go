package handlers

import (
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
)

type SendWebPushMessageHandler struct {
	ClientId string
}

func NewSendWebPushMessageHandler(clientId string) SendWebPushMessageHandler {
	return SendWebPushMessageHandler{ClientId: clientId}
}

func (sendWebPushMessageHandler SendWebPushMessageHandler) SendMessageToClient(message string, clientSubscription *webpush.Subscription) {
	fmt.Printf("Sending push notification to client %v", clientSubscription)
	// Send Notification
	resp, err := webpush.SendNotification([]byte("Push subscription successful"), clientSubscription, &webpush.Options{
		Subscriber:      "chatt-server@chatt-server.com",
		VAPIDPublicKey:  "BM221uCcUB6tJBektDBpuhrFtvECNs7mcShfG6NUnUUR1lV7vGWmWMm7eNZ0ztW4IjDPsGOAG9sQOkjP1hC_23A",
		VAPIDPrivateKey: "9LhvZAWJpanJGmkhA416muEYCWOyqzCbV_5P-Z_WR-c",
		TTL:             30,
	})
	if err != nil {
		fmt.Printf("Error sending push notification = %v", err)
	}
	defer resp.Body.Close()
}
