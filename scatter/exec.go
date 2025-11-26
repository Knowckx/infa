package scatter

import (
	"os/exec"
	"strings"

	"github.com/Knowckx/infa/util"
	"github.com/pkg/errors"
	"log/slog"
)

type Command struct {
	CmdSrc    string
	Fragments []string
	*exec.Cmd
}

func ExecCmd(in string) error {
	cmd := NewCommand(in)
	return cmd.Run()
}

func NewCommand(in string) *Command {
	out := &Command{}
	out.CmdSrc = ParseExec(in)
	out.Fragments = ParseToFragments(out.CmdSrc)
	cmdLen := len(out.Fragments)
	if cmdLen == 1 {
		out.Cmd = exec.Command(out.Fragments[0])
	} else {
		out.Cmd = exec.Command(out.Fragments[0], out.Fragments[1:cmdLen]...)
	}
	return out
}

func (t *Command) RunCommand() error {
	slog.Info("Run Command", "command", t.CmdSrc)
	err := t.Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ParseExec(in string) string {
	cmd := util.CompressSpaces(in)
	cmd = strings.TrimSpace(cmd)
	return cmd
}

func ParseToFragments(in string) []string {
	return strings.Split(in, " ")
}
