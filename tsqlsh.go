package main

import (
	"flag"
	"fmt"

	"github.com/twistingmercury/color"
	"github.com/twistingmercury/go-tsqlsh/cli"
)

var (
	dbn = flag.String("d", "master", "the name of the database to be connected to.")
	hlp = flag.String("h", "", "Help message.")
	usr = flag.String("u", "", "Connect with the user account.")
	svr = flag.String("s", "localhost", "The name of the server hosting the database.")
	pwd = flag.String("p", "", "User's password.")
	exe = flag.String("e", "", "Execute the TSQL statement and exit.")
	fln = flag.String("f", "", "Exectute the commands from a TSQL file, then exit.")
	rwr = flag.String("rw", "", "Sets the application intent to ReadWrite.")
)

func main() {
	flag.Parse()
	args := cli.ParseFlags(dbn, hlp, usr, svr, pwd, exe, fln, rwr)

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
	fmt.Println("$ tsqlsh -s <server_name> -u <user_name> -p <password> [-d <database_name>] [-e <tsql_cmd>] [-f <tsql_file>] [-rw]")
	flag.Usage()
}

func printErrors(errors []string) {
	for _, e := range errors {
		color.Red(e)
	}
}
