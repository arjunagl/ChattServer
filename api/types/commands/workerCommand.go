package commands

type Command string

const (
	SendMessage           Command = "SendMessage"
	SetClientSubscription Command = "SetClientSubscription"
)

var workerCommands = [...]string{
	"SendMessage",
	"SetClientSubscription",
}

type WorkerCommand struct {
	Command Command     `json:"command"`
	Details interface{} `json:"details"`
}
