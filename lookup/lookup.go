package lookup

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/domainr/whois"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/gosuri/uitable"
	"golang.org/x/text/language"
)

const (
	defaultLanguage = "en"
	defaultRegion   = "US"
)


// Available:
// 0 = in progress
// 1 = available
// 2 = not available
type domaincheck struct {
	Domain    string
	Available int
}

var domainExtensions = []string{"com", "net", "org", "nl", "de", "io"}

var domainscheck = []domaincheck{}

// Provider interface provides a way to build the URI
// for each provider.

// Providers tracks loaded providers.
// var Providers map[string]Provider

func init() {
	// 	Providers = make(map[string]Provider)
}

// AddProvider should be called within your provider's init() func.
// This will register the provider so it can be used.
// func AddProvider(name string, provider Provider) {
// 	Providers[name] = provider
// }

// Search builds a search URL and opens it in your browser.
func QueryHandler(p string, domain string, verbose bool) error {
	domainscheck = []domaincheck{}

	if strings.HasSuffix(domain, ".*") {
		domainscheck = multipleDomains(domain)
	} else {
		domainscheck = singleDomain(domain)
	}

	// This should be drawn first but I can't find a way to clear the table and redraw it with new information.
	// createTable(domainscheck)
	createTable(WhoisArr(domainscheck))

	return nil
}

func singleDomain(domain string) []domaincheck {
	domainscheck = append(domainscheck, domaincheck{domain, 0})

	return domainscheck
}

func multipleDomains(domain string) []domaincheck {
	domain = strings.TrimSuffix(domain, ".*")

	for _, extension := range domainExtensions {
		domainscheck = append(domainscheck, domaincheck{fmt.Sprintf("%s.%s", domain, extension), 0})
	}

	return domainscheck
}


func createTable(domainscheck []domaincheck) {
	writer := uilive.New()
	writer.Start()

	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()

	table := uitable.New()
	table.MaxColWidth = 50

	for _, domain := range domainscheck {
		indicator := ""
		if domain.Available == 0 {
			indicator = fmt.Sprintf("[%s]", white("●")) // Find better load icon
		} else if domain.Available == 1 {
			indicator = fmt.Sprintf("[%s]", green("✓"))
		} else if domain.Available == 2 {
			indicator = fmt.Sprintf("[%s]", red("×"))
		}
		table.AddRow(indicator, domain.Domain)
	}

	fmt.Fprintln(writer, table)
	writer.Stop() // flush and stop rendering
}

func WhoisArr(domains []domaincheck) []domaincheck {
	domainscheck = domains
	for i, domain := range domains {
		domainscheck[i] = domaincheck{domain.Domain, doWhois(domain.Domain, false)}
	}
	return domainscheck
}

func doWhois(qwhois string, verbose bool) int {
	request, err := whois.NewRequest(qwhois)
	FatalIf(err)

	response, err := whois.DefaultClient.Fetch(request)
	FatalIf(err)

	if verbose {
		fmt.Printf("%s\n", response)
	}

	// TODO: Split this into separate files for each extension
	// For example
	// dotcom.provider:
	// 	"No match"
	r := regexp.MustCompile(`(No match)|(^NOT FOUND)|(^Not fo|AVAILABLE)|(^No Data Fou|has not been regi|No entri)|(Status: free)|(.nl is free)`)
	if v := response.String(); r.MatchString(v) {
		return 1
	} else {
		return 2
	}
}

// Region returns the users region code.
// Eg. "US", "GB", etc
func Region() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultRegion
	}

	region, _ := tag.Region()

	return region.String()
}

// Language returns the users language code.
// Eg. "en", "es", etc
func Language() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultLanguage
	}

	base, _ := tag.Base()

	return base.String()
}

func locale() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		return ""
	}

	locale := strings.Split(lang, ".")[0]

	return locale
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
