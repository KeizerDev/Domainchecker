package main

import (
	"fmt"
	"os"
	// Load necessary domainprovider.
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotbe"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotcom"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotde"
	_ "github.com/KeizerDev/domainchecker/domainproviders/doteu"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotio"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotnet"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotnl"
	_ "github.com/KeizerDev/domainchecker/domainproviders/dotorg"

	// Load necessary providers.
	_ "github.com/KeizerDev/domainchecker/providers/godaddy"
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
