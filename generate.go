package leetgode

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
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

var _ Cmd = &GenerateCmd{}

type GenerateCmd struct{}

func (c *GenerateCmd) Name() string {
	return "generate"
}

func (g *GenerateCmd) MaxArg() int {
	return 1
}

func (g *GenerateCmd) Usage() string {
	return "Generate the skeleton code with the test file by id"
}

func (g *GenerateCmd) Run(ctx context.Context, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	cli, err := NewLeetCode()
	if err != nil {
		return err
	}
	q, err := cli.GetQuestionByID(ctx, id)
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

	// TODO: どうやってファイル保存とテストしやすさを分けようか？
	path := buildPath(q.QuestionID, q.Slug)
	fmt.Printf("save at %q\n", path)
	if err := ioutil.WriteFile(path, buf.Bytes(), 0644); err != nil {
		return err
	}
	tess, err := gotests.GenerateTests(path, nil)
	if err != nil {
		return err
	}
	if len(tess) == 0 {
		return fmt.Errorf("failed to generate test")
	}
	if err := ioutil.WriteFile(tess[0].Path, tess[0].Output, 0644); err != nil {
		return err
	}
	return nil
}
