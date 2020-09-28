package leetgode

import (
	"context"
)

var _ Cmd = &HelpCmd{}

type HelpCmd struct{}

func (c *HelpCmd) Name() string {
	return "help"
}

func (c *HelpCmd) Usage() string {
	return "Help shows usages"
}

func (c *HelpCmd) MaxArg() int {
	return 0
}

func (c *HelpCmd) Run(ctx context.Context, args []string) error {
	return nil
}
