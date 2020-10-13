package leetgode

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var _ Cmd = &ExecCmd{}

type ExecCmd struct{}

func (c *ExecCmd) Name() string {
	return "exec"
}

func (c *ExecCmd) MaxArg() int {
	return 1
}

func (c *ExecCmd) Usage() string {
	return "Submit solution"
}

// TODO: refactoring exec and test.
func (c *ExecCmd) Run(ctx context.Context, out io.Writer, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	session := os.Getenv("LEETCODE_SESSION")
	token := os.Getenv("LEETCODE_TOKEN")

	cli, err := NewLeetCode(fillAuth(session, token))
	if err != nil {
		return err
	}
	q, err := cli.GetQuestionByID(ctx, id)
	if err != nil {
		return err
	}
	fp := buildPath(q.QuestionID, q.Slug)
	code, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	tr, err := cli.Submit(ctx, q, string(code))
	if err != nil {
		return err
	}
	fmt.Fprint(out, "now sending")
	for {
		res, err := cli.Check(ctx, q, tr)
		if err != nil {
			return err
		}
		// FIXME: pretty print
		if res.State == "SUCCESS" {
			fmt.Fprintf(out, `
executed id: %s
problem title: %s
result: %s
`, q.QuestionID, q.Slug, res.StatusMsg)
			break
		} else {
			fmt.Fprint(out, ".")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
