package leetgode

import (
	"context"
	"fmt"
	"io"
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

func (c *ListCmd) Run(ctx context.Context, out io.Writer, _ []string) error {
	cli, err := NewLeetCode()
	if err != nil {
		return err
	}

	pairs, err := cli.GetStats(ctx)
	if err != nil {
		return err
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Stat.FrontendQuestionID < pairs[j].Stat.FrontendQuestionID })

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	for _, pair := range pairs {
		locked := ""
		if pair.PaidOnly {
			locked = "🔒"
		}
		fmt.Fprintf(w, "%4d\t%s\t%s\t%s\n", pair.Stat.FrontendQuestionID, pair.Stat.QuestionTitle, pair.Difficulty.Level, locked)
	}

	return w.Flush()
}
