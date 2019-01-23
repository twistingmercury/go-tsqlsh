package cli

import (
	"errors"
)

// Error text for invalid/missing args
const (
	ExecAndFileMutualExclusive = "The 'e' flag and 'f' flag are mutually exclusive."
	ServerRequired             = "The 's' flag (server) is required."
	DatabaseRequired           = "The 'd' flag (database) is required."
	UserNameRequired           = "The 'u' flag (user) is required."
	PasswordRequired           = "The 'p' flag (password) is required."
)

// CmdLineArgs represents a parsed and validated set of
// command line arguemnts from when the app was launched.
type CmdLineArgs struct {
	Database   string
	help       bool
	Username   string
	Server     string
	Password   string
	ExecuteCmd string
	Filename   string
	readWrite  bool
	Errors     []string
}

// HasErrors indicates if there were any errors in the way the cli
// arguments were passed in.
func (c *CmdLineArgs) HasErrors() bool {
	return len(c.Errors) > 0
}

// ShowHelp instructs the system to display the help for
func (c *CmdLineArgs) ShowHelp() bool {
	return c.help
}

// ReadWrite indicates if the connection should allow CRUD and DML statements.
func (c *CmdLineArgs) ReadWrite() bool {
	return c.readWrite
}

// ParseFlags validates the cli flags and returns a type of CmdLineArgs.
func ParseFlags(svr, dbn, usr, pwd, exe, fln string, hlp, rwr bool) (cla CmdLineArgs) {
	cla.help = hlp || (svr == "" && dbn == "" && usr == "" && pwd == "" && exe == "" && fln == "" && !rwr)

	if cla.help {
		return
	}

	cla.Server = svr
	cla.Database = dbn
	cla.Username = usr
	cla.Password = pwd
	cla.ExecuteCmd = exe
	cla.Filename = fln
	cla.Errors = make([]string, 0)
	cla.readWrite = rwr

	if exe != "" && fln != "" {
		cla.Errors = append(cla.Errors, ExecAndFileMutualExclusive)
	}

	if svr == "" {
		cla.Errors = append(cla.Errors, ServerRequired)
	}

	if dbn == "" {
		cla.Errors = append(cla.Errors, DatabaseRequired)
	}

	if usr == "" {
		cla.Errors = append(cla.Errors, UserNameRequired)
	}

	if pwd == "" {
		cla.Errors = append(cla.Errors, PasswordRequired)
	}
	return
}

func BuildCommand() (*string, error) {
	return nil, errors.New("not implemented")
}
