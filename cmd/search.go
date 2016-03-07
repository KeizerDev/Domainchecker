package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/KeizerDev/domainchecker/domainproviders"
	"github.com/KeizerDev/domainchecker/lookup"
	"github.com/KeizerDev/domainchecker/providers"
	"github.com/spf13/cobra"
)

const (
	appName = "domainchecker"
	version = "0.0.1"
)

// Flag variables
var displayVersion bool
var verbose bool
var provider string
var listProviders bool
var domainProviders bool
var binary string
var port int
var certpem string
var keypem string

// SearchCmd is the main command for Cobra.
var SearchCmd = &cobra.Command{
	Use:   "domainchecker <query>",
	Short: "Check domain availability from your cli",
	Long:  `Check domain availability from your cli and pass it to a domain provider. Made with ❤ by KeizerDev`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Error] %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	prepareFlags()
}

func prepareFlags() {
	SearchCmd.PersistentFlags().BoolVarP(
		&displayVersion, "version", "", false, "display version")
	SearchCmd.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false, "display url when opening")
	SearchCmd.PersistentFlags().StringVarP(
		&provider, "provider", "p", "godaddy", "set buy provider")
	SearchCmd.PersistentFlags().BoolVarP(
		&listProviders, "list-providers", "l", false, "list supported providers")
	SearchCmd.PersistentFlags().BoolVarP(
		&domainProviders, "list-extensions", "e", false, "list supported extensions")
	SearchCmd.PersistentFlags().StringVarP(
		&certpem, "cert", "c", "", "location of cert.pem for TLS")
	SearchCmd.PersistentFlags().StringVarP(
		&keypem, "key", "k", "", "location of key.pem for TLS")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if displayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if listProviders {
		fmt.Printf(providers.DisplayProviders())
		return nil
	}

	if domainProviders {
		fmt.Printf(domainproviders.DisplayProviders())
		return nil
	}

	query := strings.Join(args, " ")

	st, err := os.Stdin.Stat()
	if err != nil {
		// os.Stdin.Stat() can be unavailable on Windows.
		if runtime.GOOS != "windows" {
			return fmt.Errorf("Failed to stat Stdin: %s", err)
		}
	} else {
		if st.Mode()&os.ModeNamedPipe != 0 {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("Failed to read from Stdin: %s", err)
			}

			query = strings.TrimSpace(fmt.Sprintf("%s %s", query, bytes))
		}
	}

	if query != "" {
		err := lookup.QueryHandler(provider, query, verbose)
		// err := domainproviders.Search(binary, provider, query, verbose)
		if err != nil {
			return err
		}
	} else {
		// Don't return an error, help screen is more appropriate.
		cmd.Help()
	}

	return nil
}
