package app

import (
	"github.com/matteobovetti/redshift-tool/internal/conf"
	"github.com/matteobovetti/redshift-tool/internal/pkg/domain"
	"github.com/sirupsen/logrus"
)

var serviceLocator *ServiceLocator

type ServiceLocator struct {
	retentionConf conf.RetentionConf
	copyConf      conf.CopyConf
	unloadConf    conf.UnloadConf
	retentioner   domain.Retentioner
	copier        domain.Copier
	unloader      domain.Unloader
}

func InitServiceLocator(
	redshiftHostname string,
	redshiftPort string,
	redshiftDatabase string,
	redshiftUsername string,
	redshiftPassword string,
	dryRun bool,
	// retention command
	redshiftSchemaToRetain string,
	// copy command
	table string,
	source string,
	copyFormat string,
	copyCredentials string,
	copyOptions string,
	// unload command
	unloadQuery string,
	destination string,
	unloadFormat string,
	unloadCredentials string,
	unloadOptions string,
) {
	if serviceLocator != nil {
		panic("Service Locator already initialised")
	}

	rsConfig := conf.RedshiftConf{
		RedshiftHostname: redshiftHostname,
		RedshiftPort:     redshiftPort,
		RedshiftDatabase: redshiftDatabase,
		RedshiftUsername: redshiftUsername,
		RedshiftPassword: redshiftPassword,
	}

	serviceLocator = &ServiceLocator{
		retentionConf: conf.RetentionConf{
			RedshiftConfig:         rsConfig,
			RedshiftSchemaToRetain: redshiftSchemaToRetain,
			DryRun:                 dryRun,
		},
		copyConf: conf.CopyConf{
			RedshiftConfig: rsConfig,
			Table:          table,
			Source:         source,
			Format:         copyFormat,
			Credentials:    copyCredentials,
			Options:        copyOptions,
		},
		unloadConf: conf.UnloadConf{
			RedshiftConfig: rsConfig,
			UnloadQuery:    unloadQuery,
			Destination:    destination,
			Format:         unloadFormat,
			Credentials:    unloadCredentials,
			Options:        unloadOptions,
		},
	}
}

// GetServiceLocator expose the service locator outside this package
func GetServiceLocator() *ServiceLocator {
	if serviceLocator == nil {
		panic("Service Locator not initialised")
	}

	return serviceLocator
}

func (sl *ServiceLocator) GetRetentioner() domain.Retentioner {
	if sl.retentioner != nil {
		return sl.retentioner
	}

	sl.retentioner = domain.NewRetentioner(&sl.retentionConf)

	return sl.retentioner
}

func (sl *ServiceLocator) GetCopy() domain.Copier {
	if sl.copier != nil {
		return sl.copier
	}

	sl.copier = domain.NewCopier(&sl.copyConf)

	return sl.copier
}

func (sl *ServiceLocator) GetUnload() domain.Unloader {
	if sl.unloader != nil {
		return sl.unloader
	}

	sl.unloader = domain.NewUnloader(&sl.unloadConf)

	return sl.unloader
}

func (sl *ServiceLocator) Close() {
	logrus.Info("Service locator closed.")
}
