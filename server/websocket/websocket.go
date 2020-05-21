package websocket

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	lock sync.Mutex

	m map[string]RouterConfig
}

// RouterConfig 路由配置
type RouterConfig struct {
	Path string
	Func func(*ClientWS, map[string]string) bool
}

// ClientWS ...
type ClientWS struct {
	*websocket.Conn
	rw sync.Mutex
}

var (
	r        router
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	} // use default options
)

// Registry 路由注册
func Registry(cfg RouterConfig) {
	r.lock.Lock()
	defer r.lock.Unlock()
}

func init() {
	r = router{
		m: make(map[string]RouterConfig),
	}
}

func (c *ClientWS) write(messageType int, data []byte) error {
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
func (c *ClientWS) Write(data []byte) (n int, err error) {
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
	client := ClientWS{
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

func (c *ClientWS) decode(mt int, message []byte) bool {
	switch mt {
	case websocket.TextMessage:
		data := make(userMessage)
		json.Unmarshal(message, &data)
		if !handle(data, c) {
			return false
		}
	case websocket.BinaryMessage:
	case websocket.CloseMessage:
	case websocket.PingMessage:
	case websocket.PongMessage:
	default:
		return false
	}
	return true
}

func handle(data map[string]string, c *ClientWS) bool {
	if f, ok := r.m["method"]; ok {
		return f.Func(c, data)
	}
	return false
}
