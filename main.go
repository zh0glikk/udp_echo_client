package main

import (
	"os"

	"udp_echo_client/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
