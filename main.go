package main

import (
	"davinci-game/config"
	"davinci-game/consts"
	"davinci-game/handlers"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	app.Get("/ping", handlers.Ping)

	env := config.GetRunEnv()

	switch env {
	case consts.Production:
		certPath := os.Getenv("CERT_PATH")
		keyPath := os.Getenv("KEY_PATH")
		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
		}
		app.ListenTLS(":"+port, certPath, keyPath)
	case consts.Development:
		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
		}
		app.Listen("localhost:" + port)
	}
}
