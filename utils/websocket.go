package utils

import (
	"github.com/gofiber/websocket/v2"
	"time"
)

func CloseSocketIfAlive(conn *websocket.Conn) error {
	if conn != nil {
		err := conn.SetReadDeadline(time.Now())
		if err != nil {
			return err
		}
		if _, _, err = conn.ReadMessage(); err != nil && websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
			// connection is already closed
		} else {
			err = conn.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
