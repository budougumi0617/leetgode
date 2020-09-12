package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/budougumi/leetgode"
)

const (
	LIST = "list"
	PICK = "pick"
	GENERATE = "generate"
)

func main() {
	flag.Parse()
	sub := flag.Arg(0)
	if len(sub) == 0 {
		fmt.Printf("TODO: show help\n")
		return
	}
	cli, err := leetgode.NewLeetCode()
	if err != nil {
		fmt.Printf("failed client generation: %v\n", err)
		os.Exit(1)
	}
	ctx := context.Background()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("invalid sub command %q\n", args)
		os.Exit(1)
	}
	switch sub {
	case LIST:

	case PICK:
		q, err := cli.GetQuestion(ctx, args[1])
		if err != nil {
			fmt.Printf("failed GetQuestion(%q): %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("result: %#v\n", q)
	case GENERATE:
		if err := leetgode.GenerateCmd(ctx, args[1]);err != nil {
			fmt.Printf("failed GenerateCmd(ctx, %q): %v\n", args[1], err)
			os.Exit(1)
		}
	default:
		fmt.Printf("invalid sub command %q\n", sub)
		os.Exit(1)
	}
}
