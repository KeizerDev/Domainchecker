package dotbe

import (
	"fmt"

	"github.com/KeizerDev/domainchecker/domainproviders"
)

func init() {
	domainproviders.AddDomainProvider("be", &DomainProvider{})
}

// DomainProvider merely implements the DomainProvider interface.
type DomainProvider struct{}

// GetRegex return regex to use in whois request
func (p *DomainProvider) GetRegex() string {
	return fmt.Sprintf("(Status: AVAILABLE^)")
}

// Return used extension
func (p *DomainProvider) GetExtension() string {
	return fmt.Sprintf("be")
}
