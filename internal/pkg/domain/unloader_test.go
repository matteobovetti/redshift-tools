package domain

import (
	"strings"
	"testing"
)

func TestUnloadCommandBuild(t *testing.T) {
	sqlExpected := "UNLOAD (select * from mytable limit 100) TO s3://my-bucket/data-products/ CREDENTIALS 'aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>' FORMAT AS PARQUET;"

	sql := buildUnloadCommand(
		"select * from mytable limit 100",
		"s3://my-bucket/data-products/",
		"aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>",
		"PARQUET",
		"",
	)

	if strings.Compare(sql, sqlExpected) != 0 {
		t.Fatal("UNLOAD command is not well generated.")
	}
}
