package types

import (
	"github.com/gorilla/websocket"
)

// ClientConnection client connection structure
type ClientConnection struct {
	Connection *websocket.Conn
}

// ClientConnections existing client connections
type ClientConnections = map[string]ClientConnection
