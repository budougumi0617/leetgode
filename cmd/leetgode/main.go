package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/budougumi0617/leetgode"
)

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Args))
}

func run(stdout, stderr io.Writer, args []string) int {
	if len(args) < 2 {
		cmd := &leetgode.HelpCmd{}
		if err := cmd.Run(context.Background(), []string{}); err != nil {
			fmt.Fprintf(stderr, "help comamnd is faield: %v", err)
		}
		return 1
	}
	sub := args[1]
	f := flag.NewFlagSet(sub, flag.ContinueOnError)
	f.Usage = func() {
		if err := leetgode.ShowUsage(stdout); err != nil {
			fmt.Fprintf(stderr, "failed show useage: %v\n", err)
		}
	}
	var v bool
	f.BoolVar(&v, "v", false, "show debug print")
	if err := f.Parse(args[2:]); err == flag.ErrHelp {
		return 1
	} else if err != nil {
		fmt.Fprintf(stderr, "%s with invalid args: %v\n", sub, err)
		return 1
	}

	log.SetOutput(stderr)
	if !v {
		log.SetOutput(ioutil.Discard)
	}

	if cmd, ok := leetgode.CmdMap[leetgode.CmdName(sub)]; ok {
		args := f.Args()
		if len(args) != cmd.MaxArg() {
			fmt.Fprintf(stderr, "%s expects %d options, but %d options\n", cmd.Name(), cmd.MaxArg(), len(args))
			return 1
		}
		if err := cmd.Run(context.Background(), args); err != nil {
			fmt.Fprintf(stderr, "main: err %v", err)
			return 1
		}
	} else {
		fmt.Fprintf(stderr, "unknown command %q", sub)
		return 1
	}
	return 0
}
