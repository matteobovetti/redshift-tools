package cmd

import (
	"github.com/matteobovetti/redshift-tool/internal/app"
	"github.com/urfave/cli/v2"
)

func getUnloadCommand(baseFlags []cli.Flag, baseBeforeFunc cli.BeforeFunc) *cli.Command {

	baseFlags = append(baseFlags,
		&cli.StringFlag{
			Name:     "unload-query",
			Usage:    "Redshift query used to unload.",
			EnvVars:  []string{"RST_RS_UNLOAD_QUERY"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "destination",
			Usage:    "Destination S3 path.",
			EnvVars:  []string{"RST_RS_DESTINATION_S3_PATH"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "unload-format",
			Usage:    "Unload format. Can be CSV of PARQUET.",
			EnvVars:  []string{"RST_RS_UNLOAD_DESTINATION_FORMAT"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "unload-credentials",
			Usage:    "Data souce credentials to be used on unload command. it can be: 'aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>' or 'aws_access_key_id=<access-key-id>;aws_secret_access_key=<secret-access-key>'",
			EnvVars:  []string{"RST_RS_UNLOAD_DESTINATION_CREDENTIALS"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "unload-options",
			Usage:    "Additional options to give to the unload command.",
			EnvVars:  []string{"RST_RS_UNLOAD_OPTIONS"},
			Required: true,
		})

	cmd := &cli.Command{
		Name:   "unload",
		Usage:  "Performe a Redshift UNLOAD command.",
		Action: runUnload,
		Flags:  baseFlags,
		Before: baseBeforeFunc,
	}

	return cmd
}

func runUnload(c *cli.Context) error {
	serviceLocator := app.GetServiceLocator()

	defer serviceLocator.Close()

	serviceLocator.GetUnload().Unload()

	return nil
}
