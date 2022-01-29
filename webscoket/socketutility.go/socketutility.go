package registersocket

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var ClientIDShared []string
var PoolIDs []string
var Ids []string
var clientids []string
var Das []Data
var Check bool
var createdid string

type Client struct {
	Conn      *websocket.Conn
	Pool      *Pool
	ID        string   `json:"id,omitempty"`
	PoolID    string   `json:"poolid,omitempty"`
	Color     string   `json:"color,omitempty"`
	Points    int      `json:"points,omitempty"`
	Left      float32  `json:"left,omitempty"`
	Top       float32  `json:"top,omitempty"`
	Right     float32  `json:"right,omitempty"`
	Bottom    float32  `json:"bottom,omitempty"`
	ClientIDs []string `json:"clientids,omitempty"`
}

type Data struct {
	ID        string   `json:"id,omitempty"`
	PoolID    string   `json:"poolid,omitempty"`
	Color     string   `json:"color,omitempty"`
	Points    int      `json:"points,omitempty"`
	Left      float32  `json:"left,omitempty"`
	Top       float32  `json:"top,omitempty"`
	Right     float32  `json:"right,omitempty"`
	Bottom    float32  `json:"bottom,omitempty"`
	ClientIDs []string `json:"clientids,omitempty"`
}

type Datas struct {
	d []Data
}
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	PoolID     string
	ClientID   string
}

type Message struct {
	Type       int    `json:"type"`
	Body       string `json:"body"`
	Clientdata string `json:"client"`
	ClientID   string `json:"clientid"`
	// Clients    []Client `json:"clients"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (c *Client) Read() {
	//close client/conn
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		//take in a message
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}

		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

func GetClientID() string {
	ii := rand.Intn(100)
	createdid = fmt.Sprint(ii)
	return createdid
}
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
