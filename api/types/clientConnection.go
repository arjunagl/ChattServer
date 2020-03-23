package types

import (
	"github.com/arjunagl/ChattServer/api/types"
	"github.com/gorilla/websocket"
)

// ClientConnection client connection structure
type ClientConnection struct {
	SocketConnection  *websocket.Conn
	CientID           string
	ClientConnections types.ClientConnections
}

// ClientConnections existing client connections
type ClientConnections = map[string]ClientConnection
