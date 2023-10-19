package main

import (
	"os"

	"github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root"
)

// main is the entrypoint of the meshplayctl command-line tool
func main() {
	// Execute the root command
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
