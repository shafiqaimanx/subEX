package src

import (
	"regexp"
)

func GrabItem(domainRequested string, getRequestResult string) []string {
	var matches []string
	re := regexp.MustCompile(`url=[A-Za-z]+://[A-Za-z0-9]+\.` + regexp.QuoteMeta(domainRequested))
	matches = re.FindAllString(string(getRequestResult), -1)
	return matches
}