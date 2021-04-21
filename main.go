package main

import (
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
