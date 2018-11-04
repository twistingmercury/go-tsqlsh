package cli

// CmdLineArgs represents a parsed and validated set of
// command line arguemnts from when the app was launched.
type CmdLineArgs struct {
	Database   string
	help       bool
	Username   string
	Server     string
	Password   string
	ExecuteCmd *string
	Filename   *string
	ReadWrite  bool
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

// ParseFlags validates the cli flags and returns a type of CmdLineArgs.
func ParseFlags(dbn, hlp, usr, svr, pwd, exe, fln, rwr *string) (cla CmdLineArgs) {
	cla.help = hlp != nil

	if cla.help {
		return
	}

	cla.Database = *dbn
	cla.ExecuteCmd = exe
	cla.Filename = fln
	cla.Errors = make([]string, 0)
	cla.ReadWrite = rwr != nil

	if exe != nil && fln != nil {
		cla.Errors = append(cla.Errors, "The 'e' flag and 'f' flag are mutually exclusive.")
	}

	if svr == nil {
		cla.Errors = append(cla.Errors, "The 's' flag (server) is required.")
	}

	if usr == nil {
		cla.Errors = append(cla.Errors, "The 'u' flag (user) is required.")
	}

	if pwd == nil {
		cla.Errors = append(cla.Errors, "The 'p' flag (password) is required.")
	}
	return
}
