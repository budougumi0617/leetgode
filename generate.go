package leetgode

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/cweill/gotests"
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
package main

/**
 * {{.Content}}
**/
/**
 * {{.SampleTestCase}}
**/
{{.CodeDefinition.DefaultCode}}
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
	input := &Format{
		Referer:        q.Referer,
		Content:        q.Content,
		CodeDefinition: c,
		SampleTestCase: q.SampleTestCase,
	}
	tmpl, err := template.New("submitFile").Parse(submitFile)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.String())
	path := "tmp/hoge.go"
	if err := ioutil.WriteFile(path, buf.Bytes(), 0644); err != nil {
		return err
	}
	tess, err := gotests.GenerateTests(path, nil)
	if err != nil {
		return err
	}
	fmt.Printf("%q\n", tess[0].Path)
	if err := ioutil.WriteFile(tess[0].Path, tess[0].Output, 0644); err != nil {
		return err
	}
	return nil
}
