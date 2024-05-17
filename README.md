# Redshift tools

Tool written in Golang for manage multiple operation on top of Redshift.

Command implemented right now:
- retention
- copy
- unload

### Lead Maintainer
Matteo Bovetti - matteobovetti@gmail.com


### Common CLI parameters

Inside `.env.dist` you can find full list of env variables used in this tool.

```shell
NAME:
   AWS Redshift Tools - A new cli application

USAGE:
   AWS Redshift Tools [global options] command [command options] [arguments...]

COMMANDS:
   retention  Run a retention command on 6 months old tables.
   copy       Performe a Redshift COPY command.
   unload     Performe a Redshift UNLOAD command.
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --hostname value  Redshift hostname [$RST_RS_HOSTNAME]
   --port value      Redshift port (default: 0) [$RST_RS_PORT]
   --database value  Redshift database [$RST_RS_DATABASE]
   --username value  Redshift username [$RST_RS_USERNAME]
   --password value  Redshift password [$RST_RS_PASSWORD]
   --dry-run         dry run mode, activate for NOT apply mutation (default: false) [$RST_DRY_RUN]
   --help, -h        show help

```

Common variables:
```shell
RST_RS_HOSTNAME=
RST_RS_PORT=
RST_RS_DATABASE=
RST_RS_USERNAME=
RST_RS_PASSWORD=
RST_DRY_RUN=
```

Retention variables:
```shell
RST_RS_SCHEMA_TO_RETAIN=
```

Copy variables:
```shell
RST_RS_TABLE_TO_COPY=
RST_RS_COPY_SOURCE_S3_PATH=
RST_RS_COPY_SOURCE_FORMAT=
RST_RS_COPY_SOURCE_CREDENTIALS=
RST_RS_COPY_OPTIONS=
```

Unload variables:
```shell
RST_RS_UNLOAD_QUERY=
RST_RS_DESTINATION_S3_PATH=
RST_RS_UNLOAD_DESTINATION_FORMAT=
RST_RS_UNLOAD_DESTINATION_CREDENTIALS=
RST_RS_UNLOAD_OPTIONS=
```

### How to contribute

Build with docker:

```shell
make build
```

Run a command:

```shell
make run-<COMMAND>
```
