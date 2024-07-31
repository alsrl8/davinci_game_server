package game

import (
	"davinci-game/utils"
	"github.com/gofiber/websocket/v2"
)

type Room struct {
	ID                string
	OwnerEmail        string
	SocketConnections map[string]*websocket.Conn
}

func NewRoom(ownerEmail string) *Room {
	room := Room{
		ID:                utils.MakeRandomString(12),
		OwnerEmail:        ownerEmail,
		SocketConnections: make(map[string]*websocket.Conn),
	}
	return &room
}
