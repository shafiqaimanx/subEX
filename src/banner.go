package src

import (
	"fmt"
)

const (
	// Standard indicators
	ERROR = CRIMSON + "Error" + RESET
	WARN = ORANGE + "Warning" + RESET
	INFO = ROYALBLUE + "Info" + RESET

	// Bold indicators 
	BERROR = BOLD + CRIMSON + "Error" + RESET
	BWARN = BOLD + ORANGE + "Warning" + RESET
	BINFO = BOLD + ROYALBLUE + "Info" + RESET
)

func SubexBanner() {
	fmt.Printf("%s╔═╗╦ ╦╔╗ ╔═╗═╗╔═%s\n", MEDIUMAQUAMARINE, RESET)
	fmt.Printf("%s╚═╗║ ║╠╩╗║╣  ║╣ %s\n", MEDIUMAQUAMARINE, RESET)
	fmt.Printf("%s╚═╝╚═╝╚═╝╚═╝═╝╚═ %sv0.3.0%s\n", MEDIUMAQUAMARINE, BOLD, RESET)
}

func HelpMenu() {
	SubexBanner()
	fmt.Printf("%ssubEX - passively enumerating subdomains.%s\n\n", DARKGRAY, RESET)
	fmt.Printf("%s%sUsage%s: %ssubEX%s [flags]\n", BOLD, UNDERSCORE, RESET, BOLD, RESET)
	fmt.Printf("%s%sExample%s: subEX -d example.com\n\n", BOLD, UNDERSCORE, RESET)
	fmt.Printf("%s%sFlags%s:\n", BOLD, UNDERSCORE, RESET)
	fmt.Println("  -d   <domain>        Domain for enumerating subdomains.")
	fmt.Printf("  -i   <interaction>   Interaction of pages in search engine. (%sdefault:10%s)\n", ITALIC, RESET)
	fmt.Printf("  -o   <output>        Outputting the results into a file. (%sdefault:subex.txt%s)\n", ITALIC, RESET)
}

func DisclaimerInfo(requestedDomain string) {
	requestedDomain = MEDIUMAQUAMARINE + ITALIC + requestedDomain + RESET

	SubexBanner()
	fmt.Println("")
	fmt.Printf("%s: If you get blocked by search engine providers or any misuse.\n", BWARN)
	fmt.Printf("%s: The author assumes no responsibility!\n", BWARN)
	fmt.Printf("%s: Enumerating subdomains for %s\n", BINFO, requestedDomain)
	fmt.Printf("%s═════════════════════════════════════════════════════════════════════%s\n", MEDIUMAQUAMARINE, RESET)
}