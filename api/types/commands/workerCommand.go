package commands

type Command string

const (
	SendMessage Command = "SendMessage"
)

var workerCommands = [...]string{
	"SendMessage",
}

// func (workerCommand Command) String() string {
// 	return workerCommands[workerCommand]
// }

type WorkerCommand struct {
	// Command Command     `json:"command"`
	Command Command     `json:"command"`
	Details interface{} `json:"details"`
}
