package leetgode

import (
	"context"
	"fmt"
)

func ListCmd(ctx context.Context) error {
	cli, err := NewLeetCode()
	if err != nil {
		return err
	}

	pairs, err := cli.GetStats(ctx)
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		fmt.Printf("%d\t%s\t%s\n", pair.Stat.QuestionID, pair.Stat.QuestionTitleSlug, pair.Difficulty.Level)
	}
	return nil
}
