package helper

import (
	"fmt"
	"time"
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

func DefaultFileName(nameOfOutputFile string) string {
	dateFormat := fmt.Sprintf("%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	timeFormat := fmt.Sprintf("%d%d%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	outputFileName := fmt.Sprintf("%s_%s_%s", dateFormat, timeFormat, nameOfOutputFile)
	return string(outputFileName)
}