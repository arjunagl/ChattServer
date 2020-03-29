package types

import (
	"github.com/gorilla/websocket"
)

// ClientConnection client connection structure
type ClientConnection struct {
	SocketConnection *websocket.Conn
	CientID          string
}
