package websocket

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/wziww/bubble-test/common/docker"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
} // use default options
type clientWS struct {
	*websocket.Conn
	rw sync.Mutex
}

func (c *clientWS) write(messageType int, data []byte) error {
	c.rw.Lock()
	defer c.rw.Unlock()
	err := c.WriteMessage(messageType, data)
	if err != nil {
		logrus.Println("write:", err)
		return err
	}
	return nil
}

// Write(p []byte) (n int, err error)
func (c *clientWS) Write(data []byte) (n int, err error) {
	err = c.write(websocket.TextMessage, data)
	return len(data), err
}

// Upgrade ...
func Upgrade(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Print("upgrade:", err)
		return
	}
	defer c.Close()
	client := clientWS{
		Conn: c,
	}
	for {
		mt, message, err := client.ReadMessage()
		if err != nil {
			logrus.Println("read:", err)
			break
		}
		if !client.decode(mt, message) {
			break
		}
	}
}

type userMessage map[string]string

func (c *clientWS) decode(mt int, message []byte) bool {
	switch mt {
	case websocket.TextMessage:
		data := make(userMessage)
		json.Unmarshal(message, &data)
		resp := docker.ImagesPull(context.Background(), "docker.io/prom/alertmanager:latest")
		if resp != nil {
			io.Copy(c, resp)
		}
		c.write(websocket.TextMessage, []byte(`{code:200,message:"success"}`))
	case websocket.BinaryMessage:
	case websocket.CloseMessage:
	case websocket.PingMessage:
	case websocket.PongMessage:
	default:
		return false
	}
	return true
}
