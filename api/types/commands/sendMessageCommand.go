package commands

type SendMessageCommandDetails struct {
	Message string
	To      string
}

type SendMessageCommand struct {
	WorkerCommand
	Details SendMessageCommandDetails
}
