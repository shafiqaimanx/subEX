package helper

import (
	"os"
	"fmt"
	"regexp"
	"strconv"

	"github.com/shafiqaimanx/subex/src"
)

func DomainInputChecker(requestedDomain string) string {
	var userInputDomain string
	regexPattern := regexp.MustCompile(`^[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)
	checkPattern := regexPattern.MatchString(requestedDomain)
	if checkPattern {
		userInputDomain = requestedDomain
	} else {
		fmt.Printf("%s Invalid domain format\n", src.ERROR)
		os.Exit(0)
	}
	return userInputDomain
}

func PageInputChecker(requestedPage string) int {
	regexPattern := regexp.MustCompile(`^\d+$`)
	checkPattern := regexPattern.MatchString(requestedPage)
	if checkPattern {
		userInputPage, err := strconv.Atoi(requestedPage)
		if err != nil {
			fmt.Printf("%s %v\n", src.ERROR, err)
		}
		return userInputPage
	} else {
		fmt.Printf("%s Invalid input value\n", src.ERROR)
		os.Exit(0)
	}
	return 0
}