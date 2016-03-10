package namecheap

import (
	"fmt"
	"net/url"

	"github.com/KeizerDev/domainchecker/providers"
)

func init() {
	providers.AddProvider("namecheap", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for godaddy.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.namecheap.com/domains/registration/results.aspx?domain=%s", url.QueryEscape(q))
}
