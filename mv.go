package main

import (
	"fmt"
	"os"
//	"io"

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
			fmt.Printf("File Open Error. [%s]\n", sFromFile)
		} else	{
			fmt.Printf("File Open Success. [%s]\n", sFromFile)
		}

		iRet := moveFile(sFromFile, sToFile)

		if iRet != 1 {
			fmt.Printf("File Move Fail. [%s to %s]\n", sFromFile, sToFile)
		} else {
			fmt.Printf("File Move success. [%s to %s]\n", sFromFile, sToFile)
		}
				
	}
}

func OpenFile(sFromFile string) int {
	
 	fp, _ := os.OpenFile(sFromFile, os.O_RDONLY, os.FileMode(0644))
	if fp == nil {
		return -1
	} else {
		return 1
	}
	defer fp.Close()
	
	return 1
}

func moveFile(sFromFile string, sToFile string) int {

	link := os.Link(sFromFile , sToFile )
	if link != nil {
		return -1
	}

	unlink := os.Remove(sFromFile )
	if unlink != nil {
		return -1
	}
	
	return 1
}
