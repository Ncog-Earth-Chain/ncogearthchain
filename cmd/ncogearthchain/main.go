package main

import (
	"fmt"
	"os"

	"github.com/Ncog-Earth-Chain/ncogearthchain/cmd/ncogearthchain/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
