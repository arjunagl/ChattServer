package types

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// ClientConnection client connection structure
type ClientConnection struct {
	Connection *websocket.Conn
}

// ClientConnections existing client connections
type ClientConnections = map[uuid.UUID]ClientConnection
