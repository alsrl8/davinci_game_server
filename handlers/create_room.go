package handlers

import (
	"davinci-game/core"
	"davinci-game/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateRoomHandler(c *fiber.Ctx) error {
	email, err := utils.GetUserEmail(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to get user email from token"})
	}

	prevRoom, has := game.EmailRoomMap[email]
	if has {
		// TODO Game Room Owner가 새 연결을 시도한 경우, 기존 Room에 있던 모든 유저는 Close 해야 한다.
		conn, has := prevRoom.SocketConnections[email]
		if has {
			if conn != nil {
				err = conn.Close()
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to close previous game room"})
				}
			}
			delete(prevRoom.SocketConnections, email)
		}
	}

	newRoom := game.NewRoom(email)
	game.EmailRoomMap[email] = newRoom

	return c.JSON(fiber.Map{
		"roomId":         newRoom.ID,
		"roomOwnerEmail": newRoom.OwnerEmail,
	})
}
