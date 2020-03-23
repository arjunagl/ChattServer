package commands

type Command string

const (
	SendMessage Command = "SendMessage"
)

var workerCommands = [...]string{
	"SendMessage",
}

type WorkerCommand struct {
	Command Command     `json:"command"`
	Details interface{} `json:"details"`
}
