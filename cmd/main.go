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
	// TODO: rustのCLIはこういうフォーマット
	// ログインしていないと、解答済みかわからない。
	//✔ [  1 ] Two Sum                                                      Easy   (45.52 %)
	//[  2 ] Add Two Numbers                                              Medium (33.63 %)
	//[  3 ] Longest Substring Without Repeating Characters               Medium (30.24 %)
	//[  4 ] Median of Two Sorted Arrays                                  Hard   (29.34 %)
	//[  5 ] Longest Palindromic Substring                                Medium (29.33 %)
	//[  6 ] ZigZag Conversion                                            Medium (35.98 %)
	//✔ [  7 ] Reverse Integer                                              Easy   (25.77 %)
	case LIST:

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
