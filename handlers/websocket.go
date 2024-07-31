package handlers

import (
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
