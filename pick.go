package leetgode

import (
	"context"
	"fmt"
	"strconv"
)

var _ Cmd = &PickCmd{}

type PickCmd struct {
}

func (c *PickCmd) Name() string {
	return "pick"
}

func (c *PickCmd) Usage() string {
	return "Pick a problem by id"
}

func (c *PickCmd) MaxArg() int {
	return 1
}

func (c *PickCmd) Run(ctx context.Context, args []string) error {
	cli, err := NewLeetCode()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	q, err := cli.GetQuestionByID(ctx, id)
	if err != nil {
		return err
	}
	fmt.Printf("result: %#v\n", q)

	return nil
}
