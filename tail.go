package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var sFile string = ""

	 if ( len(os.Args) != 2 ) {
		fmt.Println("Usage : tail [FileName]")
		os.Exit(2)
	} else {		
		sFile = os.Args[1]
		openFile(sFile)
	}
}

func openFile(sFile string) int {

	var iLineCount int = 0
	var iPrintLineCount int = 0

	iLineCount = getLineCount(sFile)
	iPrintLineCount = iLineCount - 5


	file, fp := os.OpenFile(sFile, os.O_RDONLY, os.FileMode(0644))
	if fp != nil {
		fmt.Printf("File Open Fail. [%s]\n", sFile)
		return -1
	}

	reader := bufio.NewReaderSize(file, 4*1024)
	
	printTailResult(reader, iPrintLineCount)

	return 1
}

func printTailResult(reader *bufio.Reader, iPrintCount int) int{
	
	var iReadCount int = 0

	line, isPre, _ := reader.ReadLine()
	
	//for fp == nil && !isPre  {

	for !isPre  {
		s := string(line)
		iReadCount++
		if len(s) > 0  && (iReadCount > iPrintCount){
			fmt.Printf("%s\n", s )
		}
		line, isPre, _ = reader.ReadLine()
	}

	return 1
}

func getLineCount(sFile string) int {
	iLineCount := int(0)
	fp, err := os.Open(sFile)

	if err != nil {
		fmt.Printf("File Open Fail in getLineCount func. [%s]\n", sFile)
		return -1
	}

	defer fp.Close()
	scan := bufio.NewScanner(fp)
	for scan.Scan() {
		iLineCount++
	}
	return iLineCount
}









