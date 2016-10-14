package main

import (
	"fmt"
	"os"
	"io"

)

func main() {
	var sFromFile string = ""
	var sToFile string = ""

         if ( len(os.Args) < 3 || len(os.Args) > 3 ) {
		fmt.Println("Usage: cp [From File] [To File]")
		os.Exit(2)
	} else {
		sFromFile = os.Args[1]
		sToFile = os.Args[2]
		
		iReturn := OpenFile(sFromFile)

		if iReturn != 1 {
			fmt.Printf("File Open Error. [%s]", sFromFile)
		} else	{
			fmt.Printf("File Open Success. [%s]", sFromFile)
		}

		iRet := CopyFile(sFromFile, sToFile)

		if iRet != 1 {
			fmt.Printf("File Copy Fail. [%s to %s]\n", sFromFile, sToFile)
		} else {
			fmt.Printf("File Copy success. [%s to %s]\n", sFromFile, sToFile)
		}
				
	}
}

func OpenFile(sFromFile string) int {
	
 	fp, _ := os.OpenFile(sFromFile, os.O_RDONLY, os.FileMode(0644))
	if fp == nil {
		return -1
	} 
	defer fp.Close()
	
	return 1
}

func CopyFile(sFromFile string, sToFile string) int {
	
	rfp, err := os.Open(sFromFile)	
	if err != nil {
		return -1
	} 
	defer rfp.Close()

  	wfp, err := os.Create(sToFile)
	if err != nil {
		return -1
	} 
	defer wfp.Close()	

	n, err := io.Copy(wfp, rfp)
	if err != nil {
		return -1
	}
	fmt.Printf(" (ReadByte=%d)\n",n)
	
	return 1
}
