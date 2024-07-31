package routes

import (
	"davinci-game/handlers"
	"davinci-game/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", handlers.Ping)
	app.Use("/", middlewares.JWTMiddleware)
	app.Get("/", handlers.Websocket, handlers.Ws)
	app.Post("/create-room", handlers.CreateRoomHandler)
}
