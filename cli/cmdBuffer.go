package cli

// CmdBuffer is a stack-like stucture meant to
// buffer sql statements for easy re-use.
type CmdBuffer struct {
	cmds []string
}

// NewCmdBuffer creates a new CmdBuffer
func NewCmdBuffer() CmdBuffer {
	return CmdBuffer{
		cmds: make([]string, 0),
	}
}
