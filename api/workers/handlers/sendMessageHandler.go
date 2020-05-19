package handlers

import (
	"encoding/json"
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
	fmt.Printf("Outgoing %+v\n", command.Details)
	mapstructure.Decode(command.Details, &sendMessageCommand.SendMessageCommandDetails)

	// select the destination client channel
	destinationChannel := workerChannels[sendMessageCommand.To]
	destinationChannel <- command

}

func (sendMessageHandler SendMessageHandler) SendMessgeToClient(command commands.WorkerCommand, clientConnection types.ClientConnection) {
	sendMessageSocketCommand := commands.SendMessageSocketCommand{}
	mapstructure.Decode(command.Details, &sendMessageSocketCommand)
	sendMessageSocketCommand.From = sendMessageHandler.ClientId
	jsonEncodedCommand, _ := json.Marshal(sendMessageSocketCommand.SendMesasgeSocketCommandDetails)
	msg := []byte(jsonEncodedCommand)
	err := clientConnection.SocketConnection.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println(err)
	}
}
