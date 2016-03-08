package dotlol

import (
	"fmt"

	"github.com/KeizerDev/domainchecker/domainproviders"
)

func init() {
	domainproviders.AddDomainProvider("lol", &DomainProvider{})
}

// DomainProvider merely implements the DomainProvider interface.
type DomainProvider struct{}

// GetRegex return regex to use in whois request
func (p *DomainProvider) GetRegex() string {
	return fmt.Sprintf("( is available for registration.)")
}

// BuildURI generates a search URL for GitHub.
func (p *DomainProvider) GetExtension() string {
	return fmt.Sprintf("lol")
}
