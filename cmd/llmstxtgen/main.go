package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/syumai/llmstxtgen"
)

type Mode string

const (
	ModeNormal Mode = "normal"
	ModeFull   Mode = "full"
	ModeAll    Mode = "all"
)

var mode = flag.String("mode", "normal", "generate mode: normal, full, all")

func main() {
	flag.Parse()
	if *mode != string(ModeFull) {
		fmt.Fprintf(os.Stderr, "only full mode is supported")
		os.Exit(1)
	}

	reader, err := llmstxtgen.Full(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate txt: %v", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to copy: %v", err)
		os.Exit(1)
	}
}
