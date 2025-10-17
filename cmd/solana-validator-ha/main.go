package main

import (
	"os"

	"github.com/sol-strategies/solana-validator-ha/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
