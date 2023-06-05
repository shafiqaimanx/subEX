package src

import (
	"strings"
)

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