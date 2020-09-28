package leetgode

import (
	"context"
	"fmt"
)

type CmdName string

const (
	LIST     CmdName = "list"
	PICK             = "pick"
	GENERATE         = "generate"
	TEST             = "test"
	EXEC             = "exec"
)

type Cmd interface {
	Name() string
	Usage() string
	MaxArg() int
	Run(context.Context, []string) error
}

var CmdMap = map[CmdName]Cmd{
	EXEC:     &ExecCmd{},
	LIST:     &ListCmd{},
	GENERATE: &GenerateCmd{},
	TEST:     &TestCmd{},
	PICK:     &PickCmd{},
}

func buildPath(id, slug string) string {
	format := "tmp/%s.%s.go"
	return fmt.Sprintf(format, id, slug)
}
