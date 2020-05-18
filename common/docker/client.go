package docker

import (
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"github.com/wziww/bubble-test/common/config"
)

// New ...
func New() *client.Client {
	cfg := config.Get()
	c, err := client.NewClient(cfg.Docker.Host, config.DefaultAPIVersion, nil, nil)
	if err != nil {
		log.Errorln(err)
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
