package domainproviders

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/KeizerDev/domainchecker/launcher"
)

// Provider interface provides a way to build the URI
// for each provider.
type DomainProvider interface {
	GetExtension() string
	GetRegex() string
}

// Providers tracks loaded providers.
var DomainProviders map[string]DomainProvider

func init() {
	DomainProviders = make(map[string]DomainProvider)
}

// AddProvider should be called within your provider's init() func.
// This will register the provider so it can be used.
func AddDomainProvider(name string, domainProvider DomainProvider) {
	DomainProviders[name] = domainProvider
}

// Search builds a search URL and opens it in your browser.
func Search(binary string, p string, q string, verbose bool) error {
	prov, err := ExpandDomainProviders(p)
	if err != nil {
		return err
	}

	builder := DomainProviders[prov]

	if builder != nil {
		extension := builder.GetExtension()
		// regex := builder.GetRegex()
		if verbose {
			fmt.Printf("%s\n", extension)
		}

		return launcher.OpenURI(binary, extension)
	}

	return fmt.Errorf("Provider %q not supported!\n", prov)
}

// DisplayProviders displays all the loaded providers.
func DisplayProviders() string {
	names := DomainProviderNames()

	return fmt.Sprintf("%s\n", strings.Join(names, "\n"))
}

// ExpandProvider expands the passed in provider to the full value.
func ExpandDomainProviders(domainProvider string) (string, error) {
	names := DomainProviderNames()
	r := regexp.MustCompile(domainProvider + `^`)

	validProviders := []string{}
	for _, n := range names {
		// Exact match returns immediately.
		if n == domainProvider {
			return n, nil
		}

		if r.Match([]byte(n)) {
			validProviders = append(validProviders, n)
		}
	}

	switch len(validProviders) {
	case 0:
		return "", fmt.Errorf("No provider found for %q", domainProvider)
	case 1:
		return validProviders[0], nil
	default:
		return "", fmt.Errorf("Multiple providers matched %q: %v", domainProvider, validProviders)
	}
}

// ProviderNames returns a sorted slice of provider names.
func DomainProviderNames() []string {
	names := []string{}

	for key := range DomainProviders {
		names = append(names, key)
	}

	sort.Strings(names)
	return names
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
