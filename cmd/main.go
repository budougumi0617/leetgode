package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/budougumi/leetgode"
)

const (
	LIST     = "list"
	PICK     = "pick"
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
		if err := leetgode.ListCmd(ctx); err != nil {
			// TODO: ソートする！！
			fmt.Printf("failed ListCmd: %v\n", err)
			os.Exit(1)
		}
	case PICK:
		id, err := strconv.Atoi(args[1])
		if err != nil || id == 0 {
			fmt.Printf("cannot get id: %q\n", args[1])
			os.Exit(1)
		}
		q, err := cli.GetQuestionByID(ctx, id)
		if err != nil {
			fmt.Printf("failed GetQuestion(%q): %v\n", args[1], err)
			os.Exit(1)
		}
		fmt.Printf("result: %#v\n", q)
	case GENERATE:
		id, err := strconv.Atoi(args[1])
		if err != nil || id == 0 {
			fmt.Printf("cannot get id: %q\n", args[1])
			os.Exit(1)
		}
		if err := leetgode.GenerateCmd(ctx, id); err != nil {
			fmt.Printf("failed GenerateCmd(ctx, %q): %v\n", args[1], err)
			os.Exit(1)
		}
	default:
		fmt.Printf("invalid sub command %q\n", sub)
		os.Exit(1)
	}
}
