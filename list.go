package leetgode

import (
	"context"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

var _ Cmd = &ListCmd{}

type ListCmd struct{}

func (c *ListCmd) Name() string {
	return "list"
}

func (c *ListCmd) MaxArg() int {
	return 0
}

func (c *ListCmd) Usage() string {
	return "List problems"
}

func (c *ListCmd) Run(ctx context.Context, _ []string) error {
	cli, err := NewLeetCode()
	if err != nil {
		return err
	}

	pairs, err := cli.GetStats(ctx)
	if err != nil {
		return err
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Stat.QuestionID < pairs[j].Stat.QuestionID })

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for _, pair := range pairs {
		fmt.Fprintf(w, "%4d\t%s\t%s\n", pair.Stat.QuestionID, pair.Stat.QuestionTitleSlug, pair.Difficulty.Level)
	}

	return w.Flush()
}
