package server

import (
	"net/http"
	"time"

	"github.com/urfave/cli"
	"github.com/wziww/bubble-test/common/config"
	"github.com/wziww/bubble-test/common/logger"
)

// Run ...
func Run(c *cli.Context) error {
	cfg, err := config.LoadConfig(c.String("config"))
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	Server := &http.Server{
		Addr:         cfg.Server.Host,
		Handler:      routerSet(),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	logger.Info("bubble-test listening at " + cfg.Server.Host + "")
	err = Server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			logger.Debug(err.Error())
			select {
			case <-time.After(60 * time.Second):
				panic(err)
			}
		}
	}
	return err
}
