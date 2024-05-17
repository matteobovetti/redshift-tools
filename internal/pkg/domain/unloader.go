package domain

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/matteobovetti/redshift-tool/internal/conf"
	"github.com/matteobovetti/redshift-tool/internal/pkg/infrastructure/database"
	"github.com/sirupsen/logrus"
)

type Unloader interface {
	Unload()
}

type unloader struct {
	config *conf.UnloadConf
}

func NewUnloader(config *conf.UnloadConf) *unloader {
	return &unloader{
		config: config,
	}
}

func (i *unloader) Unload() {

	logrus.Info("RS hostname -> ", i.config.RedshiftConfig.RedshiftHostname)
	logrus.Info("RS port -> ", i.config.RedshiftConfig.RedshiftPort)
	logrus.Info("RS database -> ", i.config.RedshiftConfig.RedshiftDatabase)
	logrus.Info("RS username -> ", i.config.RedshiftConfig.RedshiftUsername)
	logrus.Info("Unload query -> ", i.config.UnloadQuery)
	logrus.Info("S3 Destination -> ", i.config.Destination)
	logrus.Info("S3 Credentials -> ", i.config.Credentials)
	logrus.Info("S3 Format -> ", i.config.Format)
	logrus.Info("S3 Options -> ", i.config.Options)

	db := database.CreateConnection(
		i.config.RedshiftConfig.RedshiftUsername,
		i.config.RedshiftConfig.RedshiftPassword,
		i.config.RedshiftConfig.RedshiftHostname,
		i.config.RedshiftConfig.RedshiftPort,
		i.config.RedshiftConfig.RedshiftDatabase)

	defer db.Close()

	unloadSqlCommand := buildUnloadCommand(
		i.config.UnloadQuery,
		i.config.Destination,
		i.config.Credentials,
		i.config.Format,
		i.config.Options,
	)

	logrus.Info(unloadSqlCommand)

	_, err := db.Exec(unloadSqlCommand)
	if err != nil {
		logrus.Fatal(err)
	}

}

func buildUnloadCommand(unloadQuery, destination, credentials, format, options string) string {
	return fmt.Sprintf("UNLOAD (%s) "+
		"TO %s "+
		"CREDENTIALS '%s' "+
		"FORMAT AS %s"+
		"%s;",
		unloadQuery,
		destination,
		credentials,
		format,
		options)
}
