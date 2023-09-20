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
	domainToEnumerate 	:= flag.String("d", "", "Domain to enumerate")
	numberOfPages		:= flag.String("i", "10", "Page interaction [default:10]")
	nameOfOutputFile	:= flag.String("o", "subex.txt", "Output results to file")
	flag.Usage = func() {
		fmt.Printf("%s Usage: %s -d <example.com>\n", src.INFO, os.Args[0])
		fmt.Println("Options:")
		fmt.Println("	-d DOMAIN           Domain to enumerate")
		fmt.Printf("	-i INTERACTION      Page interaction [%sdefault:10%s]\n", src.GREEN, src.RESET)
		fmt.Printf("	-o OUTPUT           Output results to file [%sdefault:subex.txt%s]\n", src.GREEN, src.RESET)
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	requestedDomain := helper.DomainInputChecker(*domainToEnumerate)
	requestedPage	:= helper.PageInputChecker(*numberOfPages)
	outputFileName 	:= helper.DefaultFileName(*nameOfOutputFile)

	fmt.Println(src.SUBEXBanner())
	fmt.Printf("%s If you get blocked by search engine providers\n", src.WARN)
	fmt.Printf("%s The author assumes no responsibility\n", src.WARN)
	fmt.Printf("%s Enumerating subdomains for (%s%s%s)\n", src.INFO, src.GREEN, requestedDomain, src.RESET)

	q1 := query.QueryDNSdumpster(requestedDomain)
	q2 := query.QueryGoogleSearch(requestedDomain, requestedPage)
		
	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
	}

	dupValuesList := append(q1, q2...)
	finalDomainList := helper.RemoveDuplicates(dupValuesList)
	for _, subDomain := range finalDomainList {
		fmt.Println(subDomain)
		_, err = file.WriteString(subDomain + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%s Done enumerating\n", src.OKAY)
	fmt.Printf("%s Output saved in (%s%s%s)\n", src.INFO, src.GREEN, outputFileName, src.RESET)
}