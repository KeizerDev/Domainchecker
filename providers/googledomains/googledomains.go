package googledomains

import (
	"fmt"
	"net/url"

	"github.com/KeizerDev/domainchecker/providers"
)

func init() {
	providers.AddProvider("googledomains", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Transip.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://domains.google.com/registrar?s=%s", url.QueryEscape(q))
}
