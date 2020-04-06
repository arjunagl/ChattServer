package handlers

import (
	"fmt"
	"log"

	"github.com/arjunagl/ChattServer/api/types"
	"github.com/arjunagl/ChattServer/api/types/commands"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type SendMessageHandler struct {
	ClientId string
}

func NewSendMessageHandler(clientId string) SendMessageHandler {
	return SendMessageHandler{ClientId: clientId}
}

func (sendMessageHandler SendMessageHandler) SendMessageToChannel(command commands.WorkerCommand, workerChannels types.WorkerChannels) {

	sendMessageCommand := commands.SendMessageCommand{}
	mapstructure.Decode(command.Details, &sendMessageCommand.Details)

	// select the destination client channel
	destinationChannel := workerChannels[sendMessageCommand.Details.To]
	destinationChannel <- command

}

func (sendMessageHandler SendMessageHandler) SendMessgeToClient(command commands.WorkerCommand, clientConnection types.ClientConnection) {
	sendMessageCommand := commands.SendMessageSocketCommand{}
	mapstructure.Decode(command.Details, &sendMessageCommand.Details)
	sendMessageCommand.Details.From = sendMessageHandler.ClientId
	fmt.Printf("Sending the following details %+v", sendMessageCommand)
	msg := []byte(sendMessageCommand.Details.Message)
	err := clientConnection.SocketConnection.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println(err)
	}
}
