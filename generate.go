package leetgode

import (
	"context"
	"fmt"
)

type Format struct {
	Referer        string
	QuestionID     string
	Content        string
	Stats          string
	CodeDefinition *Code
	SampleTestCase string
	EnableRunCode  bool
	MetaData       string
}

const submitFile = `

`

// TODO: Goの関数定義を見つける
// TODO: templateを作って文字列を作る
// TODO: ファイルとして保存する
func GenerateCmd(ctx context.Context, id string) error {
	cli, err := NewLeetCode()
	if err != nil {
		return err
	}
	q, err := cli.GetQuestion(ctx, id)
	if err != nil {
		return err
	}
	var c *Code
	for _, code := range q.CodeDefinition {
		if code.Value == "golang" {
			c = code
		}
	}
	if c == nil {
		return fmt.Errorf("not found the code for Go")
	}
	fmt.Printf("%s\n", fmt.Sprint(c.DefaultCode))
	return nil
}
