package dbaccess_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/twistingmercury/go-tsqlsh/dbaccess"
)

var (
	db      = "testdb"
	port    = 1299
	mutable = false
	host    = "10.0.0.10"
)

func Test_BuildConStr_all_values_provided(t *testing.T) {
	expected := "sqlserver://me:abc@10.0.0.10:1299?ApplicationIntent=ReadOnly&app+name=tsqlsh&database=testdb"
	actual := dbaccess.BuildConStr("me", "abc", &host, &db, &port, &mutable)
	assert.Equal(t, expected, actual)
}

func Test_BuildConStr_All_Defaults(t *testing.T) {
	expected := "sqlserver://jimmy:j3nk1nz@localhost:1433?ApplicationIntent=ReadOnly&app+name=tsqlsh&database=master"
	actual := dbaccess.BuildConStr("jimmy", "j3nk1nz", nil, nil, nil, nil)
	assert.Equal(t, expected, actual)
}
