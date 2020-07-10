package dbaccess

import (
	"database/sql"
	"fmt"
	"net/url"
)

var (
	defaultPort = 1433
	defaultDb   = "master"
	defaultHost = "localhost"
	isMutable   = false
)

/**
BuildConStr is used to create a SQL Server, URL-formatted, connection string using the supplied
values.
*/
func buildConStr(usr, pwd string, host, dbname string, port int, mutable bool) (con string) {

	appIntent := "ReadOnly"
	if mutable {
		appIntent = "ReadWrite"
	}

	query := url.Values{}
	query.Add("app name", "tsqlsh")
	query.Add("database", dbname)
	query.Add("ApplicationIntent", appIntent)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(usr, pwd),
		Host:     fmt.Sprintf("%s:%d", host, port),
		RawQuery: query.Encode(),
	}

	con = u.String()

	return
}

/**
Connect takes in the values provided and creates a URL format connection string
to the target MSSQL server
*/
func connect(usr, pwd string, host, dbname string, port int, mutable bool) (db *sql.DB, err error) {
	cstr := buildConStr(usr, pwd, host, dbname, port, mutable)

	db, err = sql.Open("sqlserver", cstr)

	return
}
