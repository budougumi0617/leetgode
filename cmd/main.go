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
)

func main() {
	flag.Parse()
	sub := flag.Arg(0)
	if len(sub) == 0 {
		fmt.Printf("TODO: show help\n")
		return
	}
	cli, err := leetgode.NewLeetCode(nil)
	if err != nil {
		fmt.Printf("failed client generation: %v\n", err)
		os.Exit(1)
	}
	ctx := context.Background()
	switch sub {
	case LIST:

	case PICK:
		args := flag.Args()
		if len(args) != 2 {
			fmt.Printf("invalid sub command %q\n", args)
			os.Exit(1)
		}
		q, err := cli.GetQuestion(ctx, args[1])
		if err != nil {
			fmt.Printf("failed GetQuestion(%q): %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("result: %#v\n", q)
	default:
		fmt.Printf("invalid sub command %q\n", sub)
		os.Exit(1)
	}
}
