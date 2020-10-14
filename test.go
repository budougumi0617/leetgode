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

var _ Cmd = &TestCmd{}

type TestCmd struct{}

func (c *TestCmd) Name() string {
	return "test"
}

func (c *TestCmd) MaxArg() int {
	return 1
}

func (c *TestCmd) Usage() string {
	return "Test solution"
}

func (c *TestCmd) Run(ctx context.Context, out io.Writer, args []string) error {
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
	q, err := cli.GetQuestionByFrontendID(ctx, id)
	if err != nil {
		return err
	}
	fp := buildPath(q.QuestionID, q.Slug)
	code, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	tr, err := cli.Test(ctx, q, string(code))
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
			fmt.Fprintf(out,`
test id: %s
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
