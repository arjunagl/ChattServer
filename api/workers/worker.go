package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/arjunagl/ChattServer/api/workers/handlers"
	"github.com/mitchellh/mapstructure"
)

type Worker interface {
	Run()
}

type WorkerImp struct {
	ClientID           string
	ClientConnection   types.ClientConnection
	CommChannel        chan commands.WorkerCommand
	WorkerChannels     types.WorkerChannels
	ClientSubscription *webpush.Subscription
}

func buildHandlers(clientId string) map[commands.Command]interface{} {
	messageHandlers := make(map[commands.Command]interface{})
	messageHandlers[commands.SendMessage] = handlers.NewSendMessageHandler(clientId)
	messageHandlers[commands.SetClientSubscription] = handlers.NewSendWebPushMessageHandler(clientId)
	return messageHandlers
}

func (w WorkerImp) Run() {
	messageHandlers := buildHandlers(w.ClientID)
	go func() {
		for {
			_, message, err := w.ClientConnection.SocketConnection.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// Parse incoming message
			incomingCommand := commands.WorkerCommand{}
			if err := json.Unmarshal(message, &incomingCommand); err != nil {
				fmt.Printf("error parsing json %v", err)
			}

			switch incomingCommand.Command {
			case commands.SendMessage:
				messageHandler := messageHandlers[commands.SendMessage].(handlers.SendMessageHandler)
				messageHandler.SendMessageToChannel(incomingCommand, w.WorkerChannels)
			}
		}
	}()
	fmt.Printf("Launched worker for %v\n", w.ClientID)

	select {
	case command := <-w.CommChannel:
		switch command.Command {
		case commands.SendMessage:
			messageHandler := messageHandlers[commands.SendMessage].(handlers.SendMessageHandler)
			messageHandler.SendMessgeToClient(command, w.ClientConnection)
		case commands.SetClientSubscription:
			webPushMessageHandler := messageHandlers[commands.SetClientSubscription].(handlers.SendWebPushMessageHandler)
			setSubscriptionCommand := commands.SetClientSubscriptionCommand{}
			mapstructure.Decode(command.Details, &setSubscriptionCommand.SetClientSubscriptionCommandDetails)
			fmt.Printf("Setting the client subscription %v", setSubscriptionCommand.SetClientSubscriptionCommandDetails.ClientSubscription)
			w.ClientSubscription = setSubscriptionCommand.SetClientSubscriptionCommandDetails.ClientSubscription

			// Just a simple test
			webPushMessageHandler.SendMessageToClient("testing", w.ClientSubscription)
		}

	}
}

func NewWorker(clientID string, clientConnection types.ClientConnection, workerChannels types.WorkerChannels) Worker {
	worker := WorkerImp{
		ClientID:         clientID,
		ClientConnection: clientConnection,
		CommChannel:      make(chan commands.WorkerCommand),
		WorkerChannels:   workerChannels,
	}
	workerChannels[clientID] = worker.CommChannel
	return worker
}

func SetClientSubscription(clientID string, clientsubscription *webpush.Subscription) {
	// Set the clientID
}
