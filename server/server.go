package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/wziww/bubble-test/common/config"
)

// Run ...
func Run(c *cli.Context) error {
	cfg, err := config.LoadConfig(c.String("config"))
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	Server := &http.Server{
		Addr:         cfg.Server.Host,
		Handler:      routerSet(),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	logrus.Infoln("bubble-test listening at " + cfg.Server.Host + "")
	err = Server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			logrus.Debugln(err)
			select {
			case <-time.After(60 * time.Second):
				panic(err)
			}
		}
	}
	return err
}
