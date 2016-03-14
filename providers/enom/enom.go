package enom

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/KeizerDev/domainchecker/providers"
)

func init() {
	providers.AddProvider("enom", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Transip.
func (p *Provider) BuildURI(q string) string {
	domain := strings.Split(q, ".")
	return fmt.Sprintf("http://www.enom.com/domainsearch/search-results.aspx?sld=%s&tld=%s&searchedDomain=%s", domain[0], domain[1], url.QueryEscape(q))
}
