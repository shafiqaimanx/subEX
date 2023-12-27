package query

import (
	"fmt"
	"sync"

	"github.com/shafiqaimanx/subEX/src"
	"github.com/shafiqaimanx/subEX/src/core"
	"github.com/shafiqaimanx/subEX/src/helper"
)

func YahooSearchCrawler(requestedDomain string, requestedPage int) []string {
	var theResults []string

	for i:=1; i <= requestedPage; i++ {
		data := fmt.Sprintf("https://search.yahoo.com/search?b=%d&p=site:%s", i*7, requestedDomain)
		theResults = append(theResults, data)
	}
	return theResults
}

func YahooSearchWorker(data string, userAgent string, domainToEnumerate string, allList *[]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	search := &core.Engine{}
	search.Get(data, userAgent)
	res := search.Body()

	matches := helper.GrabRequestedDomain(domainToEnumerate, string(res), `[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*\.`)
	mutex.Lock()
	*allList = append(*allList, matches...)
	mutex.Unlock()
}

func QueryYahooSearch(requestedDomain string, requestedPage int) []string {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var allList []string

	userAgent := src.RandomUserAgent()
	crawledData := YahooSearchCrawler(requestedDomain, requestedPage)

	for _, data := range crawledData {
		wg.Add(1)
		go YahooSearchWorker(data, userAgent, requestedDomain, &allList, &wg, &mutex)
	}
	wg.Wait()

	return allList
}