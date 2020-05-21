package docker

import (
	"github.com/docker/docker/client"
	"github.com/wziww/bubble-test/common/config"
	"github.com/wziww/bubble-test/common/logger"
)

// New ...
func New() *client.Client {
	cfg := config.Get()
	c, err := client.NewClient(cfg.Docker.Host, config.DefaultAPIVersion, nil, nil)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return c
}

func release(c *client.Client) {
	if c != nil {
		c.Close()
		return
	}
}
