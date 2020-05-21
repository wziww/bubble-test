package images

import (
	"context"
	"io"

	"github.com/wziww/bubble-test/common/docker"
	"github.com/wziww/bubble-test/common/logger"
	"github.com/wziww/bubble-test/server/websocket"
)

func pull(c *websocket.ClientWS, data map[string]string) bool {
	resp, err := docker.ImagesPull(context.Background(), data["image"])
	if err != nil {
		c.Write([]byte(err.Error()))
		c.Write([]byte(`{code:400,message:"failed"}`))
		return false
	}
	if resp != nil {
		_, e := io.Copy(c, resp)
		if e != nil {
			logger.Error(e.Error())
			return false

		}
	}
	c.Write([]byte(`{code:200,message:"success"}`))
	return true
}
func init() {
	websocket.Registry(websocket.RouterConfig{
		Path: "images/pull",
		Func: pull,
	})
}
