package handlers

import (
	"fmt"

	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/mitchellh/mapstructure"
)

func SendMessage(command commands.WorkerCommand) {

	sendMessageCommand := commands.SendMessageCommand{}
	mapstructure.Decode(command.Details, &sendMessageCommand.Details)
	fmt.Printf("Send message details %v", sendMessageCommand.Details.Message)
}
