package zegsocket

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Gamesocket(c echo.Context) error {

	pool := NewPool()
	go pool.Start()
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

	return nil
}
