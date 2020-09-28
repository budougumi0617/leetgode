package leetgode

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var _ Cmd = &TestCmd{}

type TestCmd struct{}

func (c *TestCmd) MaxArg() int {
	return 1
}

func (c *TestCmd) Usage() string {
	return "submit codes and "
}

func (c *TestCmd) Run(ctx context.Context, args []string) error {
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
	tr, err := cli.Test(ctx, q, string(code))
	if err != nil {
		return err
	}
	for {
		fmt.Print("now sending")
		res, err := cli.Check(ctx, q, tr)
		if err != nil {
			return err
		}
		if res.State == "SUCCESS" {
			fmt.Printf(`
test id: %s
test name: %s
result: %s
`, q.QuestionID, q.Slug, res.StatusMsg)
			break
		} else {
			fmt.Print(".")
		}
	}
	return nil
}
