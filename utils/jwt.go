package utils

import (
	"davinci-game/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserEmail(c *fiber.Ctx) (string, error) {
	tokenString := c.Cookies("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("no user email in token")
		}
		return userEmail, nil
	}

	return "", errors.New("can not get user email from token")
}

func GetUserEmailFromSocket(c *websocket.Conn) (string, error) {
	tokenString := c.Cookies("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("no user email in token")
		}
		return userEmail, nil
	}

	return "", errors.New("can not get user email from token")
}
