package query

import (
	"fmt"
	"sync"

	"github.com/shafiqaimanx/subEX/src"
	"github.com/shafiqaimanx/subEX/src/core"
	"github.com/shafiqaimanx/subEX/src/helper"
)

func GoogleSearchCrawler(requestedDomain string, requestedPage int) []string {
	var theResults []string
	payloadList := []string {
		"https://www.google.com/search?q=site:%s+-www&start=",
		"https://www.google.com/search?q=site:%s&start=",
	}

	for i:=0; i<len(payloadList); i++ {
		dorkList := payloadList[i]
		for p:=0; p<=requestedPage; p++ {
			url := fmt.Sprintf("%s%d", dorkList, p*10)
			result := fmt.Sprintf(url, requestedDomain)
			theResults = append(theResults, result)
		}
	}
	return theResults
}

func GoogleSearchWorker(data string, userAgent string, domainToEnumerate string, allList *[]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	search := &core.Engine{}
	search.Get(data, userAgent)
	res := search.Body()
	
	matches := helper.GrabRequestedDomain(domainToEnumerate, res, `[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*\.`)
	mutex.Lock()
	*allList = append(*allList, matches...)
	mutex.Unlock()
}

func QueryGoogleSearch(requestedDomain string, requestedPage int) []string {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var allList []string

	userAgent := src.RandomUserAgent()
	crawledData := GoogleSearchCrawler(requestedDomain, requestedPage)
	
	for _, data := range crawledData {
		wg.Add(1)
		go GoogleSearchWorker(data, userAgent, requestedDomain, &allList, &wg, &mutex)
	}
	wg.Wait()

	return allList
}
