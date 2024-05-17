package cmd

import (
	"github.com/matteobovetti/redshift-tool/internal/app"
	"github.com/urfave/cli/v2"
)

func getCopyCommand(baseFlags []cli.Flag, baseBeforeFunc cli.BeforeFunc) *cli.Command {

	baseFlags = append(baseFlags,
		&cli.StringFlag{
			Name:     "table-to-copy",
			Usage:    "Redshift table to copy. Format {schema}.{table}.",
			EnvVars:  []string{"RST_RS_TABLE_TO_COPY"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "source",
			Usage:    "Source S3 path.",
			EnvVars:  []string{"RST_RS_COPY_SOURCE_S3_PATH"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "copy-format",
			Usage:    "Data source format. Can be CSV of PARQUET.",
			EnvVars:  []string{"RST_RS_COPY_SOURCE_FORMAT"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "copy-credentials",
			Usage:    "Data souce credentials to be used on copy command. it can be: 'aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>' or 'aws_access_key_id=<access-key-id>;aws_secret_access_key=<secret-access-key>'",
			EnvVars:  []string{"RST_RS_COPY_SOURCE_CREDENTIALS"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "copy-options",
			Usage:    "Additional options to give to the copy command.",
			EnvVars:  []string{"RST_RS_COPY_OPTIONS"},
			Required: true,
		})

	cmd := &cli.Command{
		Name:   "copy",
		Usage:  "Performe a Redshift COPY command.",
		Action: runCopy,
		Flags:  baseFlags,
		Before: baseBeforeFunc,
	}

	return cmd
}

func runCopy(c *cli.Context) error {
	serviceLocator := app.GetServiceLocator()

	defer serviceLocator.Close()

	serviceLocator.GetCopy().Copy()

	return nil
}
