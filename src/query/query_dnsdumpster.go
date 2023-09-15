package query

import (
	"io"
	"fmt"
	"sync"

	"github.com/shafiqaimanx/subEX/src"
	"github.com/shafiqaimanx/subEX/src/core"
	"github.com/shafiqaimanx/subEX/src/helper"
)

func DNSdumpsterWorker(requestedDomain string, urlRequested string, requestedDataBody []byte, requestedHeaders map[string]string, wg *sync.WaitGroup, resultChan chan<- []string) {
	defer wg.Done()

	var allList []string
	resp, err := core.DoPostRequest(urlRequested, requestedDataBody, requestedHeaders)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	getRequestResult, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	matches := helper.GrabRequestedDomain(requestedDomain, string(getRequestResult), `[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*\.`)
	allList = append(allList, matches...)
	resultChan <- allList
}

func QueryDNSdumpster(requestedDomain string) []string {
	var wg sync.WaitGroup
	var theResults []string
	const URL = "https://dnsdumpster.com/"
	userAgent := src.RandomUserAgent()

	search := &core.Engine{}
	search.Get(URL, userAgent)
	bodyContent := search.Body()

	csrfmiddlewaretoken := helper.ScrapeItemInBody("input[name=csrfmiddlewaretoken]", bodyContent)
	data := []byte(fmt.Sprintf(`csrfmiddlewaretoken=%s&targetip=%s&user=free`, csrfmiddlewaretoken, requestedDomain))
	headers := map[string]string {
		"User-Agent": userAgent,
		"Cookie": search.Cookies()[0],
		"Content-Type": "application/x-www-form-urlencoded",
		"Referer": URL,
	}

	resultChan := make(chan []string)
	wg.Add(1)
	go DNSdumpsterWorker(requestedDomain, URL, data, headers, &wg, resultChan)
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