package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/Doct3rJohn/knownproject/src"
)

var allList []string
var dupValues []string
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	domainToEnumerate := flag.String("d", "", "Domain name to enumerate subdomains")
	flag.Usage = func() {
		fmt.Printf("%s Usage: %s -d <example.com>\n", src.ICON1, os.Args[0])
		fmt.Printf("%s Options:\n", src.ICON1)
		fmt.Println("	-d DOMAIN	Domain name to enumerate subdomains")
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(src.KPBanner())
	fmt.Printf("%s Enumerating subdomains for %s%s%s:\n", src.ICON1, src.YELLOW, *domainToEnumerate, src.RESET)
	fmt.Printf("%s Using [%sGOOGLE%s] engine...\n", src.ICON1, src.GREEN, src.RESET)
	userAgent := src.RandomUserAgent()
	crawledData := src.EngineToCrawl(*domainToEnumerate)

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
		theResult = fmt.Sprintf("%s %s%s%s", src.ICON2, src.YELLOW, theResult, src.RESET)
		fmt.Println(theResult)
	}
	fmt.Printf("%s Done enumerating...\n", src.ICON1)
}