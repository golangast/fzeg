/*
the * is for anything goes beyond the route text meaning /ws/ss sets up the
socket.  This is because the frontend has to send data to /ws and the
websocket replaces https:// with ws:// for websocket
*/

package main

import (
	"github.com/labstack/echo/v4"
)

//Routes is for routing
func Routes(e *echo.Echo) {
	e.GET("/ws*", Gamesocket)
	e.GET("/ss", Message)
}
