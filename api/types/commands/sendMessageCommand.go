package commands

type SendMessageCommandDetails struct {
	Message string
	To      string
}

type SendMessageCommand struct {
	WorkerCommand
	Details SendMessageCommandDetails `json:"details"`
}

type SendMesasgeSocketCommandDetails struct {
	SendMessageCommandDetails
	From string
}

type SendMessageSocketCommand struct {
	WorkerCommand,
	Details SendMesasgeSocketCommandDetails `json:"details"`
}
