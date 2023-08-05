package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/Doct3rJohn/subEX/src"
)

var allList []string
var dupValues []string
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	domainToEnumerate := flag.String("d", "", "Domain to enumerate")
	numberOfPages := flag.String("i", "10", "Page interaction [default:10]")
	flag.Usage = func() {
		fmt.Printf("%s Usage: %s -d <example.com>\n", src.INFO, os.Args[0])
		fmt.Println("Options:")
		fmt.Println("	-d DOMAIN           Domain to enumerate")
		fmt.Printf("	-i INTERACTION      Page interaction [%sdefault:10%s]\n", src.GREEN, src.RESET)
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	requestedDomain := src.DomainInputChecker(*domainToEnumerate)
	requestedPage := src.PageInputChecker(*numberOfPages)
	fmt.Println(requestedDomain)
	fmt.Println(requestedPage)

	fmt.Println(src.SUBEXBanner())
	fmt.Printf("%s If you get blocked by search engine providers\n", src.WARN)
	fmt.Printf("%s The author assumes no responsibility\n", src.WARN)
	fmt.Printf("%s Enumerating subdomains for (%s%s%s)\n", src.INFO, src.GREEN, requestedDomain, src.RESET)
	
	userAgent := src.RandomUserAgent()
	crawledData := src.EngineToCrawl(requestedDomain, requestedPage)

	for _, data := range crawledData {
		wg.Add(1)
		go func(data string) {
			defer wg.Done()
			res := src.DoGetRequest(data, userAgent)
			matches := src.GrabItem(*domainToEnumerate, res)
			mutex.Lock()
			allList = append(allList, matches...)
			dupValues = src.RemoveDuplicates(allList)
			mutex.Unlock()
		}(data)	
	}
	wg.Wait()

	for i:=0; i<len(dupValues); i++ {
		theResult := src.CleanSomeItem(dupValues[i])
		fmt.Println(theResult)
	}
	fmt.Printf("%s Done enumerating\n", src.OKAY)
}