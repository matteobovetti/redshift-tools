package domain

import (
	"strings"
	"testing"
)

func TestCopyCommandBuild(t *testing.T) {
	sqlExpected := "COPY myschema.test  FROM 's3://my-bucket/data-products/' CREDENTIALS 'aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>' FORMAT AS PARQUET;"

	sql := buildCopyCommand(
		"myschema.test",
		"s3://my-bucket/data-products/",
		"aws_iam_role=<arn-aws-iam-role-MyRedshiftRole>",
		"PARQUET",
		"",
	)

	if strings.Compare(sql, sqlExpected) != 0 {
		t.Fatal("COPY command is not well generated.")
	}
}
