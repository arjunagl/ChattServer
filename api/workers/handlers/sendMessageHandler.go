package handlers

import (
	"fmt"

	"github.com/arjunagl/ChattServer/api/types/commands"
)

func SendMessage(command commands.WorkerCommand) {
	sendMessageCommand := commands.SendMessageCommand{}
	// sendMessageCommand.Command = command
	fmt.Printf("Incoming command %v", sendMessageCommand.Details.Message)
	// if err := json.Unmarshal(message, &incomingCommand); err != nil {
	// 	fmt.Printf("error parsing json %v", err)
	// }
}
