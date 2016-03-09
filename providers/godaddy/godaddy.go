package godaddy

import (
	"fmt"
	"net/url"

	"github.com/KeizerDev/domainchecker/providers"
)

func init() {
	providers.AddProvider("godaddy", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for godaddy.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://godaddy.com/domains/searchresults.aspx?checkAvail=1&domainToCheck=%s", url.QueryEscape(q))
}
