package dotcom

import (
	"fmt"

	"github.com/KeizerDev/domainchecker/domainproviders"
)

func init() {
	domainproviders.AddDomainProvider("org", &DomainProvider{})
}

// DomainProvider merely implements the DomainProvider interface.
type DomainProvider struct{}

// GetRegex return regex to use in whois request
func (p *DomainProvider) GetRegex() string {
	return fmt.Sprintf("(NOT FOUND)")
}

// Return used extension
func (p *DomainProvider) GetExtension() string {
	return fmt.Sprintf("org")
}
