package query

import (
	"sync"

	"github.com/shafiqaimanx/subEX/src"
	"github.com/shafiqaimanx/subEX/src/core"
	"github.com/shafiqaimanx/subEX/src/helper"
)

func CrtshWorker(requestedDomain string, userAgent string, domainInUsed string, wg *sync.WaitGroup, resultChan chan<- []string) {
	defer wg.Done()
	var allList []string

	search := &core.Engine{}
	search.Get(domainInUsed+requestedDomain, userAgent)
	res := search.Body()

	matches := helper.GrabRequestedDomain(requestedDomain, res, `[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*\.`)
	allList = append(allList, matches...)
	resultChan <- allList
}

func QueryCrtsh(requestedDomain string) []string {
	var wg sync.WaitGroup
	var theResults []string
	const URL = "https://crt.sh/?q="
	userAgent := src.RandomUserAgent()

	resultChan := make(chan []string)
	wg.Add(1)
	go CrtshWorker(requestedDomain, userAgent, URL, &wg, resultChan)
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for theResult := range resultChan {
		if theResult != nil {
			theResults = append(theResults, theResult...)
		}
	}
	return theResults
}