package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/shafiqaimanx/subex/src"
	"github.com/PuerkitoBio/goquery"
)

func GrabRequestedDomain(requestedDomain string, getRequestResult string, regexString string) []string {
	var matches []string
	re := regexp.MustCompile(regexString + regexp.QuoteMeta(requestedDomain))
	matches = re.FindAllString(string(getRequestResult), -1)
	return matches
}

func ScrapeItemInBody(itemObject string, bodyContent string) string {
	reader := strings.NewReader(bodyContent)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
	}

	value, exists := doc.Find(itemObject).Attr("value")
	if !exists {
		fmt.Printf("%s [ScrapeItemInBody] Item did not exists!\n", src.ERROR)
	}
	return value
}