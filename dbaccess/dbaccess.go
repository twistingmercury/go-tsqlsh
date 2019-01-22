package dbaccess

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
)

var (
	defaultPort = 1433
	defaultDb   = "master"
	defaultHost = "localhost"
)

func BuildConStr(usr, pwd string, host, dbname *string, port *int, mutable *bool) (con string) {
	if port == nil || *port == 0 {
		port = &defaultPort
	}

	if dbname == nil || len(*dbname) == 0 {
		dbname = &defaultDb
	}

	if host == nil || len(*host) == 0 {
		host = &defaultHost
	}

	appIntent := "ReadOnly"
	if mutable != nil && *mutable {
		appIntent = "ReadWrite"
	}

	query := url.Values{}
	query.Add("app name", "tsqlsh")
	query.Add("database", *dbname)
	query.Add("ApplicationIntent", appIntent)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(usr, pwd),
		Host:     fmt.Sprintf("%s:%d", *host, *port),
		RawQuery: query.Encode(),
	}

	con = u.String()

	return
}

/**
Connect takes in the values provided and creates a URL format connection string
to the target MSSQL server
*/
func Connect(usr, pwd string, host, dbname *string, port *int, mutable *bool) (db *sql.DB, err error) {
	cstr := BuildConStr(usr, pwd, host, dbname, port, mutable)

	db, err = sql.Open("sqlserver", cstr)

	return
}
