package leetgode

import (
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
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

var desc = `
Usage: leetgode is leetcode cli for gophers.

SubCommands:
`

func ShowUsage(w io.Writer) error {
	cms := make([]string, len(CmdMap))
	i := 0
	for k := range CmdMap {
		cms[i] = string(k)
		i++
	}
	sort.Strings(cms)

	tw := tabwriter.NewWriter(w, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "%s", desc)
	for _, k := range cms {
		cn := CmdName(k)
		fmt.Fprintf(tw, "\t%s\t%s\n", CmdMap[cn].Name(), CmdMap[cn].Usage())
	}
	return tw.Flush()
}

func (c *HelpCmd) Run(ctx context.Context, args []string) error {
	return ShowUsage(os.Stdout)
}
