package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shafiqaimanx/subEX/src"
	"github.com/shafiqaimanx/subEX/src/helper"
	"github.com/shafiqaimanx/subEX/src/query"
)

func main() {
	domainToEnumerate 	:= flag.String("d", "", "")
	numberOfPages		:= flag.String("i", "10", "")
	nameOfOutputFile	:= flag.String("o", "subex.txt", "")
	flag.Usage = func() {		
		src.HelpMenu()
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	requestedDomain := helper.DomainInputChecker(*domainToEnumerate)
	requestedPage	:= helper.PageInputChecker(*numberOfPages)
	outputFileName 	:= helper.DefaultFileName(*nameOfOutputFile)

	src.DisclaimerInfo(requestedDomain)
	q1 := query.QueryDNSdumpster(requestedDomain)
	q2 := query.QueryGoogleSearch(requestedDomain, requestedPage)
	q3 := query.QueryCrtsh(requestedDomain)
	q4 := query.QueryYahooSearch(requestedDomain, requestedPage)
			
	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
	}

	dupValuesList := append(append(append(q1, q2...), q3...), q4...)
	finalDomainList := helper.RemoveDuplicates(dupValuesList)
	for _, subDomain := range finalDomainList {
		fmt.Println(subDomain)
		_, err = file.WriteString(subDomain + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%s═════════════════════════════════════════════════════════════════════%s\n", src.MEDIUMAQUAMARINE, src.RESET)
	fmt.Printf("%s: Done enumerating!\n", src.BINFO)
	fmt.Printf("%s: Output saved in %s%s%s%s\n", src.BINFO, src.MEDIUMAQUAMARINE, src.ITALIC, outputFileName, src.RESET)
}