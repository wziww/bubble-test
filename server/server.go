package server

import (
	"errors"

	"github.com/urfave/cli"
	"github.com/wziww/bubble-test/common/config"
)

// Run ...
func Run(c *cli.Context) error {
	config.LoadConfig(c.String("config"))
	return errors.New("asd")
}
