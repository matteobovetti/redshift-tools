package domain

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/matteobovetti/redshift-tool/internal/conf"
	"github.com/matteobovetti/redshift-tool/internal/pkg/infrastructure/database"
	"github.com/sirupsen/logrus"
)

type Copier interface {
	Copy()
}

type copier struct {
	config *conf.CopyConf
}

func NewCopier(config *conf.CopyConf) *copier {
	return &copier{
		config: config,
	}
}

func (i *copier) Copy() {
	logrus.Info("RS hostname -> ", i.config.RedshiftConfig.RedshiftHostname)
	logrus.Info("RS port -> ", i.config.RedshiftConfig.RedshiftPort)
	logrus.Info("RS database -> ", i.config.RedshiftConfig.RedshiftDatabase)
	logrus.Info("RS username -> ", i.config.RedshiftConfig.RedshiftUsername)
	logrus.Info("Table to copy -> ", i.config.Table)
	logrus.Info("Data source to copy -> ", i.config.Source)
	logrus.Info("Data source format-> ", i.config.Format)
	logrus.Info("Data credentials-> ", i.config.Credentials)
	logrus.Info("Optional additional options-> ", i.config.Options)

	db := database.CreateConnection(
		i.config.RedshiftConfig.RedshiftUsername,
		i.config.RedshiftConfig.RedshiftPassword,
		i.config.RedshiftConfig.RedshiftHostname,
		i.config.RedshiftConfig.RedshiftPort,
		i.config.RedshiftConfig.RedshiftDatabase)

	defer db.Close()

	copySqlCommand := buildCopyCommand(
		i.config.Table,
		i.config.Source,
		i.config.Credentials,
		i.config.Format,
		i.config.Options,
	)

	logrus.Info(copySqlCommand)

	_, err := db.Exec(copySqlCommand)
	if err != nil {
		logrus.Fatal(err)
	}

}

func buildCopyCommand(table, source, credentials, format, options string) string {
	return fmt.Sprintf("COPY "+
		"%s "+
		"%s "+
		"FROM '%s' "+
		"CREDENTIALS '%s' "+
		"FORMAT AS %s"+
		"%s;",
		table,
		"",
		source,
		credentials,
		format,
		options)
}
