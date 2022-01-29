package mainesocket

import (
	"fmt"
	"net/http"

	. "github.com/golangast/fzeg/websocket/registerscoket"
	"github.com/gorilla/websocket"
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

	client := &websockets.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
