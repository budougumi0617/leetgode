package leetgode

import (
	"context"
	"fmt"
	"io"
)

type CmdName string

const (
	LIST     CmdName = "list"
	PICK             = "pick"
	GENERATE         = "generate"
	TEST             = "test"
	EXEC             = "exec"
	HELP             = "help"
)

type Cmd interface {
	Name() string
	Usage() string
	MaxArg() int
	Run(context.Context, io.Writer, []string) error
}

var CmdMap = map[CmdName]Cmd{
	EXEC:     &ExecCmd{},
	LIST:     &ListCmd{},
	GENERATE: &GenerateCmd{},
	TEST:     &TestCmd{},
	PICK:     &PickCmd{},
	HELP:     &HelpCmd{},
}

func buildPath(id int, slug string) string {
	// TODO: changeable directory
	format := "%d.%s.go"
	return fmt.Sprintf(format, id, slug)
}
