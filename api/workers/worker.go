package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/arjunagl/ChattServer/api/workers/handlers"
)

type Worker interface {
	Run()
}

type worker struct {
	ClientConnection types.ClientConnection
}

func (w worker) Run() {
	go func() {
		for {
			_, message, err := w.ClientConnection.SocketConnection.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(message))

			// Parse incoming message
			incomingCommand := commands.WorkerCommand{}
			if err := json.Unmarshal(message, &incomingCommand); err != nil {
				fmt.Printf("error parsing json %v", err)
			}
			switch incomingCommand.Command {
			case commands.SendMessage:
				handlers.SendMessage(incomingCommand)
			}
		}
	}()
	fmt.Printf("Launched worker for %v", w.ClientConnection.CientID)
}

func NewWorker(clientConnection types.ClientConnection) Worker {
	return worker{
		ClientConnection: clientConnection,
	}
}
