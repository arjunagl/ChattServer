package commands

type SendMessageCommandDetails struct {
	message string
	to      string
}

type SendMessageCommand struct {
	WorkerCommand
	details SendMessageCommandDetails
}
