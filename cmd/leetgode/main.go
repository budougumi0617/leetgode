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
	if err := run(os.Stdout, os.Stderr, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(stdout, stderr io.Writer, args []string) error {
	if len(args) < 2 {
		cmd := &leetgode.HelpCmd{}
		if err := cmd.Run(context.Background(), []string{}); err != nil {
			return fmt.Errorf("help command is failed: %w", err)
		}
		return nil
	}
	sub := args[1]
	// implement into each sub command if use different options by each sub command
	f := flag.NewFlagSet(sub, flag.ContinueOnError)
	f.Usage = func() {
		if err := leetgode.ShowUsage(stdout); err != nil {
			fmt.Fprintf(stderr, "failed show useage: %w\n", err)
		}
	}
	var v bool
	f.BoolVar(&v, "v", false, "show debug print")
	if err := f.Parse(args[2:]); err == flag.ErrHelp {
		return err
	} else if err != nil {
		return fmt.Errorf("%q command with invalid args(%q): %w", sub, args[2:], err)
	}

	log.SetOutput(stderr)
	if !v {
		log.SetOutput(ioutil.Discard)
	}

	if cmd, ok := leetgode.CmdMap[leetgode.CmdName(sub)]; ok {
		args := f.Args()
		if len(args) != cmd.MaxArg() {
			return fmt.Errorf("%q command expects %d options, but %d options\n", cmd.Name(), cmd.MaxArg(), len(args))
		}
		if err := cmd.Run(context.Background(), args); err != nil {
			return fmt.Errorf("%q command failed: %w", sub, err)
		}
	} else {
		return fmt.Errorf("unknown command %q", sub)
	}
	return nil
}
