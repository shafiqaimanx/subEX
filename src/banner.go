package src

import (
	"fmt"
)

const (
    RESET  	= "\033[0m"
    RED    	= "\033[38;5;196m"
	BLUE	= "\033[38;5;39m"
	GREEN	= "\033[38;5;48m"
    YELLOW 	= "\033[38;5;226m"
	OKAY  	= "[" + "\033[38;5;48m" 	+ "OKAY" + "\033[0m" + "]"
	WARN  	= "[" + "\033[38;5;179m" 	+ "WARN" + "\033[0m" + "]"
	INFO   	= "[" + "\033[38;5;74m" 	+ "INFO" + "\033[0m" + "]"
	ERROR  	= "[" + "\033[38;5;203m" 	+ "ERRO" + "\033[0m" + "]"
	RESULT  = "[" + "\033[38;5;213m" 	+ "RSLT" + "\033[0m" + "]"
)

//func SUBEXBanner() string {
//	fmt.Println("     	  _            ")
//	fmt.Println(" ___ _ _ | |_  ___ __  		subex - CLI tool for finding subdomain")
//   	fmt.Println("<_-<| | || . \\/ ._>\\ \\/		> _shafiqaiman_")
//   	fmt.Println("/__/`___||___/\\___./\\_\\ 	> v0.2.0")
//	return ""
//}


func SUBEXBanner() string {
	fmt.Println("          _    ___ __  _ ")
	fmt.Println(" ___ _ _ | |_ | __>\\ \\/  	subex - CLI tool for finding subdomain")
   	fmt.Println("<_-<| | || . \\| _>  \\ \\  	> _shafiqaiman_")
   	fmt.Println("/__/`___||___/|___>_/\\_\\ 	> v0.2.0")
	   return ""
}