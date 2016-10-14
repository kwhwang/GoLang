package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	var sSearchString string = ""
	var sFile string = ""

         if ( len(os.Args) == 2 ) {

		sSearchString = os.Args[1]
		grepFromStdin(sSearchString)

	} else if ( len(os.Args) == 3 ) {

		sSearchString = os.Args[1]
		sFile = os.Args[2]	
		grepFromFile(sSearchString, sFile)

	} else {		
		fmt.Println("Usage(from file)    : grep [SearchString] [File]")
		fmt.Println("         (from stdin) : (STDIN) | grep [SearchString]")
		os.Exit(2)			
	}
}

func grepFromFile(sSearchString string, sFile string) int {
	
 	file, fp := os.OpenFile(sFile, os.O_RDONLY, os.FileMode(0644))
	if fp != nil {
		fmt.Printf("File Open Fail. [%s]\n", sFile)
		return -1
	}

	reader := bufio.NewReaderSize(file, 4*1024)
	printGrepResult(reader, sSearchString )
		
	return 1
}

func grepFromStdin(sSearchString string) int {
	
	reader := bufio.NewReader(os.Stdin)
	printGrepResult(reader, sSearchString)
	return 1

}

func printGrepResult(reader *bufio.Reader, sSearchString string) int{

	line, isPre, fp := reader.ReadLine()
	
	for fp == nil && !isPre  {
		s := string(line)
		if len(s) > 0 && strings.Contains(s, sSearchString) {
			fmt.Printf("%s\n", s )
		}
		line, isPre, fp = reader.ReadLine()
	}
	
	return 1
}
