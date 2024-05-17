package conf

type RedshiftConf struct {
	RedshiftHostname string
	RedshiftPort     string
	RedshiftDatabase string
	RedshiftUsername string
	RedshiftPassword string
}

type RetentionConf struct {
	RedshiftConfig         RedshiftConf
	RedshiftSchemaToRetain string
	DryRun                 bool
}

type CopyConf struct {
	RedshiftConfig RedshiftConf
	Table          string
	Source         string
	Format         string
	Credentials    string
	Options        string
}

type UnloadConf struct {
	RedshiftConfig RedshiftConf
	UnloadQuery    string
	Destination    string
	Format         string
	Credentials    string
	Options        string
}
