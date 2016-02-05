package main

import (
	"fmt"
	"os"

	// Load necessary providers.
	_ "github.com/KeizerDev/domainchecker/providers/transip"

	"github.com/KeizerDev/domainchecker/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.SearchCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
