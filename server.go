package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"bufio"
	"io"
)

type SPCOLLECT_INFO struct {
	sAddr string
	sTime string
	sLastRecvTime string
	sReplaceCommand string
	iCount	int
	iFd	int
}

type SPCOLLECT_CONF struct {
	iPort	int
	iEOFFlag	int
	iMode	int
	iRemoveTime	int
	sLogFileName	string
	sHomeDir	string
	sFileTime	string
	sFileExt	string
	sIP_Key	string
	sIP_Map	string
	sIP_Map2	string
	iIP_Key_Size	int
	iIP_Map_Size int
	iThreadNum	int	
	iQSize		int
	iReadCount	int
	iSleepTime	int
	iFlushUse	int
	iEventCount	int
	iRemoveNewLine	int
	sDelimiter	string
	sCommand	string
	iCommandUse int
}

var gpspcollect_info = new(SPCOLLECT_CONF)

func  CheckError(err error) {
	
	if err != nil {
		fmt.Println("Error: " , err )
		os.Exit(0)
	}
}

func ReadConf() int {
	file, fp := os.OpenFile(
		"C:\\spcollect.conf",
		os.O_RDONLY,
		os.FileMode(0644))

	if fp != nil {
		fmt.Println("##")
		fmt.Println(fp)
		return -1
	}
	
	SetDefaultValue()


	read := bufio.NewReaderSize(file, 4*1024)
	line, isPre, fp := read.ReadLine()
	
	for fp == nil && !isPre {
		s := string(line)
		if strings.HasPrefix(s, "#") != true && len(s) > 0 {
			fmt.Printf("Data=[%s]\n", s)
		}
		
		line, isPre, fp = read.ReadLine()
					
	}

	if isPre {
		fmt.Println("###")
		return 0
	}

	if fp != io.EOF {
		fmt.Println(fp)
		return 0
	}	

	return 1
}

func SetDefaultValue () {
	gpspcollect_info.iPort = 514	
	gpspcollect_info.iEOFFlag = 1	
	gpspcollect_info.iMode = 0	
	gpspcollect_info.iRemoveTime	= 0
	gpspcollect_info.sHomeDir		="/ESM/spcollect"
	gpspcollect_info.sLogFileName	="spcollect.log" 
	gpspcollect_info.sFileTime		="yyyymmddhh"
	gpspcollect_info.sFileExt		="log"
	gpspcollect_info.sIP_Key		= ""
	gpspcollect_info.sIP_Map		= ""
	gpspcollect_info.sIP_Map2		= ""
	gpspcollect_info.iIP_Key_Size	= 1
	gpspcollect_info.iIP_Map_Size 	= 1
	gpspcollect_info.iThreadNum		= 1
	gpspcollect_info.iQSize		= 20000
	gpspcollect_info.iReadCount		= 5000
	gpspcollect_info.iSleepTime		= 100
	gpspcollect_info.iFlushUse 		= 1
	gpspcollect_info.iEventCount	 	= 1
	gpspcollect_info.iRemoveNewLine 	= 1
	gpspcollect_info.sDelimiter 		= "`"	
	gpspcollect_info.sCommand		= ""
	gpspcollect_info.iCommandUse 	= 0

	fmt.Println(gpspcollect_info.iPort)
	fmt.Println(gpspcollect_info.iEOFFlag )
	fmt.Println(gpspcollect_info.iMode )
	fmt.Println(gpspcollect_info.iRemoveTime)
	fmt.Println(gpspcollect_info.sHomeDir)
	fmt.Println(gpspcollect_info.sLogFileName)
	fmt.Println(gpspcollect_info.sFileTime)
	fmt.Println(gpspcollect_info.sFileExt)
	fmt.Println(gpspcollect_info.iThreadNum)
	fmt.Println(gpspcollect_info.iQSize)
	fmt.Println(gpspcollect_info.iReadCount)
	fmt.Println(gpspcollect_info.iSleepTime)
	fmt.Println(gpspcollect_info.iFlushUse )
	fmt.Println(gpspcollect_info.iEventCount)
	fmt.Println(gpspcollect_info.iRemoveNewLine )
	fmt.Println(gpspcollect_info.sDelimiter )
	fmt.Println(gpspcollect_info.sCommand)
	fmt.Println(gpspcollect_info.iCommandUse )
}

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp", ":9009")

	CheckError(err)

	Ret := ReadConf()

	//fmt.Printf("Ret=[%d]\n",Ret)	
	
	if Ret != 1 {
		fmt.Println("fail to open spcollect.conf")
		return 
	} else
	{
		fmt.Println("success to open spcollect.conf")
	}
		
	
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ", addr )

		if err != nil {
			fmt.Println("Error: ", err)
		}				

		
  		file, fp := os.OpenFile(
				"C:\\syslog.log", 
				os.O_CREATE|os.O_RDONLY|os.O_APPEND,
				os.FileMode(0644))
		if fp != nil {
			fmt.Println(fp)
			return 
		}
		
		defer file.Close()
		
		n, tmp := file.Write([]byte(buf))

		if tmp != nil {
			fmt.Println(tmp)
			return
		}
		

		
		fmt.Println("Save")
	}
}

