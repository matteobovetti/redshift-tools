package cmd

import (
	"github.com/matteobovetti/redshift-tool/internal/app"
	"github.com/urfave/cli/v2"
)

func getRetentionerCommand(baseFlags []cli.Flag, baseBeforeFunc cli.BeforeFunc) *cli.Command {

	baseFlags = append(baseFlags,
		&cli.StringFlag{
			Name:     "schema-to-retain",
			Usage:    "Redshift schema to check for retention",
			EnvVars:  []string{"RST_RS_SCHEMA_TO_RETAIN"},
			Required: true,
		})

	cmd := &cli.Command{
		Name:   "retention",
		Usage:  "Run a retention command on 6 months old tables.",
		Action: runRetentioner,
		Flags:  baseFlags,
		Before: baseBeforeFunc,
	}

	return cmd
}

func runRetentioner(c *cli.Context) error {
	serviceLocator := app.GetServiceLocator()

	defer serviceLocator.Close()

	serviceLocator.GetRetentioner().Retention()

	return nil
}
