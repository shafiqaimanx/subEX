package src

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func DomainInputChecker(requestedDomain string) string {
	var userInputDomain string
	regexPattern := regexp.MustCompile(`^[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)
	checkPattern := regexPattern.MatchString(requestedDomain)
	if checkPattern {
		userInputDomain = requestedDomain
	} else {
		fmt.Printf("%s Invalid domain format\n", ERROR)
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
			fmt.Printf("%s %v\n", ERROR, err)
		}
		return userInputPage
	} else {
		fmt.Printf("%s Invalid input value\n", ERROR)
		os.Exit(0)
	}
	return 0
}

func GrabItem(requestedDomain string, getRequestResult string) []string {
	var matches []string
	re := regexp.MustCompile(`url=[A-Za-z]+://[A-Za-z0-9]+\.` + regexp.QuoteMeta(requestedDomain))
	matches = re.FindAllString(string(getRequestResult), -1)
	return matches
}

func RemoveDuplicates(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func CleanSomeItem(dupValues string) string {
	var theResult string
	valueOne := strings.Replace(dupValues, "url=https://", "", 1)
	theResult = strings.Replace(valueOne, "url=http://", "", 1)
	return theResult
}