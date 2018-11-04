package main

import (
	"flag"
	"fmt"

	"github.com/twistingmercury/color"

	"github.com/twistingmercury/go-tsqlsh/cli"
)

var (
	dbn = flag.String("d", "", "the name of the database to be connected to.")
	svr = flag.String("s", "", "The name of the server hosting the database.")
	usr = flag.String("u", "", "The user account.")
	pwd = flag.String("p", "", "User's password.")
	exe = flag.String("e", "", "Execute the TSQL statement and exit.")
	fln = flag.String("f", "", "Exectute the commands from a TSQL file, then exit.")
	rwr = flag.Bool("rw", false, "Sets the application intent to ReadWrite.")
	hlp = flag.Bool("h", false, "Help message.")
)

func main() {
	flag.Parse()
	args := cli.ParseFlags(*svr, *dbn, *usr, *pwd, *exe, *fln, *hlp, *rwr)

	switch {
	case args.ShowHelp():
		showUsage()
	case args.HasErrors():
		printErrors(args.Errors)
		showUsage()
	}
}

// ShowUsage displays how tsqlsh should be called.
func showUsage() {
	fmt.Println("Command:")
	color.Green("  $ tsqlsh -s <server_name> -u <user_name> -p <password> [-d <database_name>] [-e <tsql_cmd>] [-f <tsql_file>] [-rw]")
	flag.Usage()
}

func printErrors(errors []string) {
	for _, e := range errors {
		color.Red("- %s", e)
	}
	fmt.Println()
}
