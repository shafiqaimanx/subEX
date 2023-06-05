package src

import (
	"fmt"
)

func EngineToCrawl(domainRequested string) []string {
	var res []string
	google := []string {
		"https://www.google.com/search?q=site:%s+-www&start=",
		"https://www.google.com/search?q=site:%s&start=",
	}

	for i:=0; i<len(google); i++ {
		dorkList := google[i]
		for p:=0; p<=9; p++ {
			url := fmt.Sprintf("%s%d", dorkList, p*10)
			result := fmt.Sprintf(url, domainRequested)
			res = append(res, result)
		}
	}
	return res
}