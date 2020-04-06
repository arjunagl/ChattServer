package types

import "github.com/arjunagl/ChattServer/api/types/commands"

// type WorkerChannel struct {
// 	CientID           string
// 	ClientConnections chan commands.WorkerCommand
// }

type WorkerChannels = map[string]chan commands.WorkerCommand
