package template

import "strings"

func getDomainShort(domain string) string {
	domainSplitted := strings.Split(domain, " ")
	short := ""

	for _, word := range domainSplitted {
		short += strings.ToLower(string(word[0]))
	}

	return short
}
