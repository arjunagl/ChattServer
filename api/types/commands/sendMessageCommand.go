package commands

type SendMessageCommandDetails struct {
	Message string
	To      string
}

type SendMessageCommand struct {
	WorkerCommand
	SendMessageCommandDetails `mapstructure:",squash"`
}

type SendMesasgeSocketCommandDetails struct {
	SendMessageCommandDetails `mapstructure:",squash"`
	From                      string
}

type SendMessageSocketCommand struct {
	WorkerCommand
	SendMesasgeSocketCommandDetails `mapstructure:",squash"`
}
