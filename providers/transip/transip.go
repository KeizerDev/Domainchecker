package transip

import (
	"fmt"
	"net/url"

	"github.com/KeizerDev/domainchecker/providers"
)

func init() {
	providers.AddProvider("transip", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Transip.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://github.com/search?utf8=✓&q=%s", url.QueryEscape(q))
}
