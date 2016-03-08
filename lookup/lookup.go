package lookup

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/KeizerDev/domainchecker/domainproviders"
	"github.com/domainr/whois"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/gosuri/uitable"
)

// Available:
// 0 = in progress
// 1 = available
// 2 = not available
type domaincheck struct {
	Domain    string
	Available int
}

var domainscheck = []domaincheck{}

func init() {
}

// Search builds a search URL and opens it in your browser.
func QueryHandler(p string, domain string, verbose bool) error {
	domainscheck = []domaincheck{}

	if strings.HasSuffix(domain, ".*") {
		domainscheck = multipleDomains(domain)
	} else {
		domainscheck = singleDomain(domain)
	}

	writer := uilive.New()
	writer.Start()
	// This should be drawn first but I can't find a way to clear the table and redraw it with new information.

	createTable(writer, domainscheck)
	createTable(writer, WhoisArr(domainscheck))

	writer.Stop() // flush and stop rendering

	return nil
}

func singleDomain(domain string) []domaincheck {
	domainscheck = append(domainscheck, domaincheck{domain, 0})

	return domainscheck
}

func multipleDomains(domain string) []domaincheck {
	domain = strings.TrimSuffix(domain, ".*")

	for _, domainext := range domainproviders.DomainProviders {
		domainscheck = append(domainscheck, domaincheck{fmt.Sprintf("%s.%s", domain, domainext.GetExtension()), 0})
	}

	return domainscheck
}

func createTable(writer *uilive.Writer, domainscheck []domaincheck) {


	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()

	table := uitable.New()
	table.MaxColWidth = 50

	for _, domain := range domainscheck {
		indicator := ""
		switch domain.Available {
		case 0:
			indicator = fmt.Sprintf("[%s]", white("●")) // Find better load icon
		case 1:
			indicator = fmt.Sprintf("[%s]", green("✓"))
		case 2:
			indicator = fmt.Sprintf("[%s]", red("×"))
		}
		table.AddRow(indicator, domain.Domain)
	}

	fmt.Fprintln(writer, table)
}

func WhoisArr(domains []domaincheck) []domaincheck {
	domainscheck = domains
	for i, domain := range domains {
		domainscheck[i] = domaincheck{domain.Domain, doWhois(domain.Domain, false)}
	}
	return domainscheck
}

func RegexBuilder() string {
	regex := ``
	for _, domainext := range domainproviders.DomainProviders {
		regex = regex + domainext.GetRegex() + `|`
	}

	// Return and just add some more bonus regex :)
	return regex + `(^Not fo|AVAILABLE)|(^No Data Fou|has not been regi|No entri)`
}

func doWhois(qwhois string, verbose bool) int {
	request, err := whois.NewRequest(qwhois)
	FatalIf(err)

	response, err := whois.DefaultClient.Fetch(request)
	FatalIf(err)

	if verbose {
		fmt.Printf("%s\n", response)
	}

	r := regexp.MustCompile(RegexBuilder())
	if v := response.String(); r.MatchString(v) {
		return 1
	} else {
		return 2
	}
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
