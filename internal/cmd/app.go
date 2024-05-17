package cmd

import (
	iapp "github.com/matteobovetti/redshift-tool/internal/app"
	"github.com/matteobovetti/redshift-tool/internal/pkg/infrastructure/log"
	"github.com/urfave/cli/v2"
)

var version string = "v1"

func GetApp() *cli.App {
	app := cli.NewApp()

	app.Version = version
	app.Name = "AWS Redshift Tools"
	app.Usage = ""
	app.HideVersion = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "hostname",
			Usage:    "Redshift hostname",
			EnvVars:  []string{"RST_RS_HOSTNAME"},
			Required: true,
		},
		&cli.IntFlag{
			Name:     "port",
			Usage:    "Redshift port",
			EnvVars:  []string{"RST_RS_PORT"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "database",
			Usage:    "Redshift database",
			EnvVars:  []string{"RST_RS_DATABASE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "username",
			Usage:    "Redshift username",
			EnvVars:  []string{"RST_RS_USERNAME"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "Redshift password",
			EnvVars:  []string{"RST_RS_PASSWORD"},
			Required: true,
		},
		&cli.BoolFlag{
			Name:     "dry-run",
			Usage:    "dry run mode, activate for NOT apply mutation",
			EnvVars:  []string{"RST_DRY_RUN"},
			Required: true,
		},
	}

	app.Before = func(c *cli.Context) error {
		return log.SetUp()
	}

	var setUp cli.BeforeFunc = func(c *cli.Context) error {
		iapp.InitServiceLocator(
			c.String("hostname"),
			c.String("port"),
			c.String("database"),
			c.String("username"),
			c.String("password"),
			c.Bool("dry-run"),
			c.String("schema-to-retain"),
			c.String("table-to-copy"),
			c.String("source"),
			c.String("copy-format"),
			c.String("copy-credentials"),
			c.String("copy-options"),
			c.String("unload-query"),
			c.String("destination"),
			c.String("unload-format"),
			c.String("unload-credentials"),
			c.String("unload-options"),
		)

		return nil
	}

	app.Commands = []*cli.Command{
		getRetentionerCommand(app.Flags, setUp),
		getCopyCommand(app.Flags, setUp),
		getUnloadCommand(app.Flags, setUp),
	}

	return app
}
