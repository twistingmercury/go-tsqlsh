package dbaccess

import (
	"database/sql"
	"errors"
)

/**
QueryResult contains the results of a sql query.
*/
type QueryResult struct {
	cols []string
	vals [][]string
}

/**
Columns contains the names of the columns that were returned.
*/
func (q *QueryResult) Columns() []string {
	return q.cols
}

/**
RowCount returns the number of rows that were returned.
*/
func (q *QueryResult) RowCount() int {
	return len(q.vals)
}

/**
Row returns a specific row within the result.
*/
func (q *QueryResult) Row(r int) []string {
	return q.vals[r]
}

/**
RowValue returns a specific value for a column within a row.
*/
func (q *QueryResult) RowValue(r, v int) string {
	return q.vals[r][v]
}

/**
MsSqlDbAccess describes an interface by which an MSSQL Server database
can be queried.
*/
type IMsSqlDbAccess interface {
	ExecQuery(query string) (*QueryResult, error)
	ExecCommand(cmd string) error
	Database() string
	Host() string
	Port() int
	Mutable() bool
	User() string
	Close() error
	State() string
}

/**
New returns a new instance of the type IMsSqlDbAccess using the supplied connection values.
*/
func New(usr, pwd string, host, dbname *string, port *int, mutable *bool) (db IMsSqlDbAccess, err error) {

	if port == nil || *port == 0 {
		port = &defaultPort
	}

	if dbname == nil || len(*dbname) == 0 {
		dbname = &defaultDb
	}

	if host == nil || len(*host) == 0 {
		host = &defaultHost
	}

	isMutable = false
	if mutable == nil {
		mutable = &isMutable
	}

	con, err := connect(usr, pwd, *host, *dbname, *port, *mutable)

	if err != nil {
		return
	}

	db = &sqlDbAccess{
		host:       *host,
		db:         *dbname,
		port:       *port,
		mutable:    *mutable,
		user:       usr,
		connection: con,
	}

	return
}

type sqlDbAccess struct {
	user       string
	pwd        string
	host       string
	db         string
	port       int
	mutable    bool
	appintent  string
	connection *sql.DB
}

func (s *sqlDbAccess) ExecQuery(query string) (*QueryResult, error) {
	return nil, errors.New("not implemented")
}
func (s *sqlDbAccess) ExecCommand(cmd string) error {
	return errors.New("not implemented")
}
func (s *sqlDbAccess) connect(usr, pwd string, host, dbname *string, port *int, mutable *bool) error {
	return errors.New("not implemented")
}
func (s *sqlDbAccess) Mutable() bool {
	return s.mutable
}
func (s *sqlDbAccess) Database() string {
	return s.db
}
func (s *sqlDbAccess) Host() string {
	return s.host
}
func (s *sqlDbAccess) Port() int {
	return s.port
}
func (s *sqlDbAccess) ApplicationIntent() string {
	return s.appintent
}
func (s *sqlDbAccess) User() string {
	return s.user
}

func (s *sqlDbAccess) State() string {
	if s.connection != nil {
		return "open"
	} else {
		return "closed"
	}
}

func (s *sqlDbAccess) Close() error {
	if s.connection == nil {
		return nil
	}
	err := s.connection.Close()
	if err != nil {
		return err
	}
	s.connection = nil
	return nil
}
