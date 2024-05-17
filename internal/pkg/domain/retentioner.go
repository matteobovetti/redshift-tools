package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/matteobovetti/redshift-tool/internal/conf"
	"github.com/matteobovetti/redshift-tool/internal/pkg/infrastructure/database"
	"github.com/sirupsen/logrus"
	"time"
)

type Retentioner interface {
	Retention()
}

type retentioner struct {
	config *conf.RetentionConf
}

func NewRetentioner(config *conf.RetentionConf) *retentioner {
	return &retentioner{
		config: config,
	}
}

type TableToRetain struct {
	SchemaName   string
	TableName    string
	CreationTime time.Time
	TableOwner   string
}

func (i *retentioner) Retention() {
	logrus.Info("RS hostname -> ", i.config.RedshiftConfig.RedshiftHostname)
	logrus.Info("RS port -> ", i.config.RedshiftConfig.RedshiftPort)
	logrus.Info("RS database -> ", i.config.RedshiftConfig.RedshiftDatabase)
	logrus.Info("RS username -> ", i.config.RedshiftConfig.RedshiftUsername)
	logrus.Info("RS schema to retain -> ", i.config.RedshiftSchemaToRetain)
	logrus.Info("RS dryrun -> ", i.config.DryRun)

	db := database.CreateConnection(
		i.config.RedshiftConfig.RedshiftUsername,
		i.config.RedshiftConfig.RedshiftPassword,
		i.config.RedshiftConfig.RedshiftHostname,
		i.config.RedshiftConfig.RedshiftPort,
		i.config.RedshiftConfig.RedshiftDatabase)

	defer db.Close()

	err := applyRetention(db, i.config.RedshiftSchemaToRetain, i.config.DryRun)
	if err != nil {
		logrus.Fatal(err)
	}

}

func applyRetention(db *sql.DB, schemaToRetain string, dryRun bool) error {

	// Extract table with 180 days old.
	retentionedTablesSql := fmt.Sprintf("SELECT "+
		"TRIM(nspname) AS schema_name, "+
		"TRIM(relname) AS table_name, "+
		"relcreationtime AS creation_time, "+
		"pg_get_userbyid(relowner) AS table_owner "+
		"FROM pg_class_info "+
		"LEFT JOIN pg_namespace ON pg_class_info.relnamespace = pg_namespace.oid "+
		"WHERE reltype != 0 "+
		"AND TRIM(nspname) = '%s' "+
		"AND relcreationtime IS NOT NULL "+
		"AND creation_time < dateadd(day, -180, current_date)"+
		"ORDER BY relcreationtime;", schemaToRetain)

	rows, err := db.Query(retentionedTablesSql)
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	logrus.Info(fmt.Sprintf("Starting to process tables in retention for schema [%s]...", schemaToRetain))

	// For every table we apply a retention command (DROP TABLE).
	for rows.Next() {
		var table TableToRetain

		err := rows.Scan(&table.SchemaName, &table.TableName, &table.CreationTime, &table.TableOwner)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info(
			fmt.Sprintf(
				"Table -> [%s], with creation time [%s] and owner [%s]",
				table.TableName,
				table.CreationTime.Format("2006-02-01"),
				table.TableOwner,
			),
		)

		if !dryRun {
			dropTable(db, &table)
		}
	}

	// Handle errors
	err = rows.Err()
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}

func dropTable(db *sql.DB, table *TableToRetain) {
	dropSql := fmt.Sprintf("DROP TABLE %s.%s;", table.SchemaName, table.TableName)

	_, err := db.Exec(dropSql)
	if err != nil {
		logrus.Fatal(err)
	}
}
