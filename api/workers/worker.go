package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/arjunagl/ChattServer/api/workers/handlers"
)

type Worker interface {
	Run()
}

type WorkerImp struct {
	ClientConnection   types.ClientConnection
	CommChannel        chan commands.WorkerCommand
	WorkerChannels     types.WorkerChannels
	ClientSubscription *webpush.Subscription
}

func buildHandlers(clientId string) map[commands.Command]interface{} {
	messageHandlers := make(map[commands.Command]interface{})
	messageHandlers[commands.SendMessage] = handlers.NewSendMessageHandler(clientId)
	return messageHandlers
}

func (w WorkerImp) Run() {
	mesasgeHandlers := buildHandlers(w.ClientConnection.CientID)
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
				messageHandler := mesasgeHandlers[commands.SendMessage].(handlers.SendMessageHandler)
				messageHandler.SendMessageToChannel(incomingCommand, w.WorkerChannels)
			}
		}
	}()
	fmt.Printf("Launched worker for %v\n", w.ClientConnection.CientID)

	select {
	case command := <-w.CommChannel:
		messageHandler := mesasgeHandlers[commands.SendMessage].(handlers.SendMessageHandler)
		messageHandler.SendMessgeToClient(command, w.ClientConnection)
	}
}

func NewWorker(clientConnection types.ClientConnection, workerChannels types.WorkerChannels) Worker {
	worker := WorkerImp{
		ClientConnection: clientConnection,
		CommChannel:      make(chan commands.WorkerCommand),
		WorkerChannels:   workerChannels,
	}
	workerChannels[clientConnection.CientID] = worker.CommChannel
	return worker
}
