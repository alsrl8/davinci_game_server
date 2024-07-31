package main

import (
	"davinci-game/config"
	"davinci-game/consts"
	"davinci-game/middlewares"
	"davinci-game/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(middlewares.NewCORS())

	routes.SetupRoutes(app)

	runByEnv(app)
}

func runByEnv(app *fiber.App) {
	env := config.GetRunEnv()
	switch env {
	case consts.Production:
		certPath := os.Getenv("CERT_PATH")
		keyPath := os.Getenv("KEY_PATH")
		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
		}
		err := app.ListenTLS(":"+port, certPath, keyPath)
		fmt.Println(err)
	case consts.Development:
		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
		}
		err := app.Listen("localhost:" + port)
		fmt.Println(err)
	}
}
