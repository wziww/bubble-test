package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/wziww/bubble-test/common/logger"
	"github.com/wziww/bubble-test/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			panic(r)
		}
	}()

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = func(c *cli.Context) error {
		return server.Run(c)
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
