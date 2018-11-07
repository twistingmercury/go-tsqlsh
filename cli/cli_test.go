package cli_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twistingmercury/go-tsqlsh/cli"
)

const (
	dbn   = "master"
	svr   = "server"
	usr   = "user"
	pwd   = "password"
	exe   = "SELECT * FROM SYS.DATABASES;"
	fln   = "t-sql.sql"
	empty = ""
)

func Test_ParseFlags_ShowHelp_Returns_True(t *testing.T) {
	args := cli.ParseFlags(empty, empty, empty, empty, empty, empty, true, false)
	assert.True(t, args.ShowHelp())
	assert.False(t, args.HasErrors())
}

func Test_ParseFlags_no_e_or_f_flags_no_errors(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, pwd, empty, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.False(t, args.HasErrors())
	assert.Equal(t, svr, args.Server)
	assert.Equal(t, dbn, args.Database)
	assert.Equal(t, usr, args.Username)
	assert.Equal(t, pwd, args.Password)
	assert.Equal(t, "", args.ExecuteCmd)
	assert.Equal(t, "", args.Filename)
	assert.Len(t, args.Errors, 0)
}

func Test_ParseFlags_e_flag_no_errors(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, pwd, exe, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.False(t, args.HasErrors())
	assert.Equal(t, svr, args.Server)
	assert.Equal(t, dbn, args.Database)
	assert.Equal(t, usr, args.Username)
	assert.Equal(t, pwd, args.Password)
	assert.Equal(t, exe, args.ExecuteCmd)
	assert.Equal(t, "", args.Filename)
	assert.Len(t, args.Errors, 0)
}

func Test_ParseFlags_f_flag_no_errors(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, pwd, empty, fln, false, false)
	assert.False(t, args.ShowHelp())
	assert.False(t, args.HasErrors())
	assert.Equal(t, svr, args.Server)
	assert.Equal(t, dbn, args.Database)
	assert.Equal(t, usr, args.Username)
	assert.Equal(t, pwd, args.Password)
	assert.Equal(t, "", args.ExecuteCmd)
	assert.Equal(t, fln, args.Filename)
	assert.Len(t, args.Errors, 0)
}

func Test_ParseFlags_e_and_f_together_flags_creates_error_msg(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, pwd, exe, fln, false, false)
	assert.False(t, args.ShowHelp())
	assert.True(t, args.HasErrors())
	assert.False(t, args.ReadWrite())
	assert.Len(t, args.Errors, 1)
	assert.Equal(t, cli.ExecAndFileMutualExclusive, args.Errors[0])
}

func Test_ParseFlags_empty_s_flag_creates_error_msg(t *testing.T) {
	args := cli.ParseFlags(empty, dbn, usr, pwd, empty, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.True(t, args.HasErrors())
	assert.False(t, args.ReadWrite())
	assert.Len(t, args.Errors, 1)
	assert.Equal(t, cli.ServerRequired, args.Errors[0])
}

func Test_ParseFlags_empty_d_flag_creates_error_msg(t *testing.T) {
	args := cli.ParseFlags(svr, empty, usr, pwd, empty, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.True(t, args.HasErrors())
	assert.False(t, args.ReadWrite())
	assert.Len(t, args.Errors, 1)
	assert.Equal(t, cli.DatabaseRequired, args.Errors[0])
}

func Test_ParseFlags_empty_u_flag_creates_error_msg(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, empty, pwd, empty, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.True(t, args.HasErrors())
	assert.False(t, args.ReadWrite())
	assert.Len(t, args.Errors, 1)
	assert.Equal(t, cli.UserNameRequired, args.Errors[0])
}

func Test_ParseFlags_empty_p_flag_creates_error_msg(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, empty, empty, empty, false, false)
	assert.False(t, args.ShowHelp())
	assert.True(t, args.HasErrors())
	assert.False(t, args.ReadWrite())
	assert.Len(t, args.Errors, 1)
	assert.Equal(t, cli.PasswordRequired, args.Errors[0])
}

func Test_ParseFlags_rw_flag_ReadWrite_is_true(t *testing.T) {
	args := cli.ParseFlags(svr, dbn, usr, pwd, empty, empty, false, true)
	assert.False(t, args.ShowHelp())
	assert.False(t, args.HasErrors())
	assert.True(t, args.ReadWrite())
	assert.Equal(t, svr, args.Server)
	assert.Equal(t, dbn, args.Database)
	assert.Equal(t, usr, args.Username)
	assert.Equal(t, pwd, args.Password)
	assert.Equal(t, "", args.ExecuteCmd)
	assert.Equal(t, "", args.Filename)
	assert.Len(t, args.Errors, 0)
}

func Test_BuildCommand_exit_command(t *testing.T) {

}
