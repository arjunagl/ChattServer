package types

import (
	"github.com/gorilla/websocket"
)

// ClientConnection client connection structure
type ClientConnection struct {
	SocketConnection  *websocket.Conn
	CientID           string
	ClientConnections ClientConnections
}

// ClientConnections existing client connections
type ClientConnections = map[string]ClientConnection
