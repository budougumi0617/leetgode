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
	flag.Parse()
	sub := flag.Arg(0)
	if len(sub) == 0 {
		fmt.Printf("TODO: show help\n")
		return
	}

	if cmd, ok := leetgode.CmdMap[leetgode.CmdName(sub)]; ok {
		if err := cmd.Run(context.Background(), flag.Args()[1:]); err != nil {
			log.Printf("main: err %v", err)
			os.Exit(1)
		}
	}
}
