package cmd

import (
	"log"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/techquest-tech/go-utils/str"
	"github.com/urfave/cli/v2"
)

//BuildVersion version for the app
var BuildVersion = "development"

//GitCommit golang version.
var GitCommit = "unknown"

func init() {
	str.ReplaceByEnv("APP_VERSION", &BuildVersion)
	logrus.Info("App version:", BuildVersion)
}

//Version print App version details.
func Version() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Action: func(c *cli.Context) error {
			log.Print("BuildVersion:\t", BuildVersion)
			log.Print("GitCommit:\t", GitCommit)
			log.Print("APP version from ENV:\t", os.Getenv("APP_VERSION"))
			return nil
		},
	}
}

//Cleanup func be called when app close.
type Cleanup func()

//CloseOnlyNotified make app close only when notify
func CloseOnlyNotified(c ...Cleanup) {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	select {
	case <-sigCh:
		logrus.Infof("Got Interrupt, app existing...")
		for _, cc := range c {
			cc()
		}
	}
}
