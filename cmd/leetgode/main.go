package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/budougumi/leetgode"
)

func main() {
	flag.Usage = func() {
		if err := leetgode.ShowUsage(os.Stdout); err != nil {
			panic(err)
		}
	}
	flag.Parse()
	sub := flag.Arg(0)
	if len(sub) == 0 {
		fmt.Printf("TODO: show help\n")
		return
	}

	if cmd, ok := leetgode.CmdMap[leetgode.CmdName(sub)]; ok {
		args := flag.Args()[1:]
		if len(args) != cmd.MaxArg() {
			fmt.Printf("%s expects %d options, but %d options\n", cmd.Name(), cmd.MaxArg(), len(args))
			os.Exit(1)
		}
		if err := cmd.Run(context.Background(), args); err != nil {
			log.Printf("main: err %v", err)
			os.Exit(1)
		}
	} else {
		log.Printf("unknown command %q", sub)
		os.Exit(1)
	}
}
