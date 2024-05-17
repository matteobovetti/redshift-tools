package main

import (
	"github.com/sirupsen/logrus"
	"os"

	command "github.com/matteobovetti/redshift-tool/internal/cmd"
)

func main() {
	app := command.GetApp()

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Error(err)
	}
}
