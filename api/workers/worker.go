package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
)

type Worker interface {
	Run()
}

type worker struct {
	ClientConnection types.ClientConnection
}

func (w worker) Run() {
	for {
		_, message, err := w.ClientConnection.SocketConnection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Message received")
		fmt.Println(string(message))

		// Parse incoming message
		incomingCommand := &commands.WorkerCommand{}
		if err := json.Unmarshal(message, incomingCommand); err != nil {
			fmt.Println(err)
		}
		fmt.Println("successfully parsed")
		fmt.Println(incomingCommand.Command)
		switch incomingCommand.Command {
		case commands.SendMessage:
			fmt.Print("Have to send the message")

		}
	}
}

func NewWorker(clientConnection types.ClientConnection) Worker {
	return worker{
		ClientConnection: clientConnection,
	}
}
