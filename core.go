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
	Usage() string
	MaxArg() int
	Run(context.Context, []string) error
}

var cmdMap = map[CmdName]Cmd{
	EXEC: &GenerateCmd{},
}

func buildPath(id, slug string) string {
	format := "tmp/%s.%s.go"
	return fmt.Sprintf(format, id, slug)
}