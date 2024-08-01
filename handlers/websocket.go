package handlers

import (
	game "davinci-game/core"
	"davinci-game/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"sync"
)

func Websocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}

var mutex sync.Mutex

var Ws = websocket.New(func(c *websocket.Conn) {

	mutex.Lock()
	email, _ := utils.GetUserEmailFromSocket(c)
	room, has := game.EmailRoomMap[email]
	if has {
		room.SocketConnections[email] = c
	}
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		err := c.Close()
		if err != nil {
			fmt.Printf("close connection error: %v", err)
		}
		mutex.Unlock()
	}()

})
