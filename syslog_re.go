package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"
)

type SPCOLLECT_CONF struct {
	sPort		string
	sFileName	string
}

var gpspcollect_info = new(SPCOLLECT_CONF)


const UDP_SOCKET_BUFFER_SIZE = 83886080
const RECV_BUFFER_SIZE = 10240



func main() {
	var i int	
	
	sPort := SetDefaultValue()
	
	ServerAddr,err := net.ResolveUDPAddr("udp", sPort)						
	ServerConn, err := net.ListenUDP("udp", ServerAddr)						
	CheckError(err)
	defer ServerConn.Close()

	err = ServerConn.SetReadBuffer(UDP_SOCKET_BUFFER_SIZE)								// 시스템 UDP Buffer 사이즈 증가

	sRecvBuf := make([]byte, RECV_BUFFER_SIZE)											
	
	sOldTime := convertStringtoTime(gpspcollect_info.sFileName)							
	fFile := openWriteFile(sOldTime)

	for {
		sNowTime := convertStringtoTime(gpspcollect_info.sFileName)			

 		n,addr,err := ServerConn.ReadFromUDP(sRecvBuf)						

		if err != nil {														
			fmt.Println("Error: ", err)
		}
		
		iReturn := compareOldNowTime(sOldTime, sNowTime) 								//Write중인 파일 명의 갱신이 필요한지 확인 (분단위가 바뀌었는지 확인) 

		if ( iReturn != 1 ) {												 
			fmt.Printf("File Switching. Old=[%s], Now=[%s]\n", sOldTime, sNowTime)
			defer fFile.Close()												
			fFile = openWriteFile(sNowTime)
			sOldTime = fmt.Sprintf("%s", sNowTime)
		} 
		
		sWriteTime  := convertStringtoTime("yyyymmdd hhmiss")				

		sDataBuf := fmt.Sprintf("[%s][%s][%s]\n",addr.IP, sWriteTime, string(sRecvBuf[0:n]))		
		
		n, tmp := fFile.Write([]byte(sDataBuf))													

		i++
		if ( i > 100 ) {																		
			fFile.Sync()
			i = 0
		}
		
		if tmp != nil {
			fmt.Println(tmp)
			return
		}
		
	}
}

func  CheckError(err error) {
	
	if err != nil {
		fmt.Println("Error: " , err )
		os.Exit(0)
	}
}

func SetDefaultValue () (string) {

	if( len(os.Args) == 3 ) {		//	[syslog port filename]
		gpspcollect_info.sPort = fmt.Sprintf(":%s", os.Args[1])
		gpspcollect_info.sFileName = fmt.Sprintf("%s", os.Args[2])
	} else {
		fmt.Printf("Usage : ./syslog [PORT] [FILENAME]")
		os.Exit(2)
	}

	fmt.Printf("PORT     = [%s]\n", gpspcollect_info.sPort)
	fmt.Printf("FILENAME = [%s]\n",gpspcollect_info.sFileName)

	return gpspcollect_info.sPort
}

func convertStringtoTime (sTime string) string {					//매개변수에 yyyymmddhhmiss 중 어느 string이 들어와도 해당하는 시간 값으로 변경해주는 함수
	var sTemp string
	var sDate string
	t := time.Now()
	sTemp = fmt.Sprintf("%s", sTime)
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "yyyy", "2006", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "yy", "06", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "mm", "01", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "dd", "02", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "hh", "15", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "mi", "04", -1))
	sTemp = fmt.Sprintf("%s", strings.Replace(sTemp, "ss", "05", -1))

	sDate = fmt.Sprintf(t.Format(sTemp))

	return sDate
}

func compareOldNowTime(sOldTime string, sNowTime string) int {		
	
	if ( strings.Contains(sNowTime, sOldTime) ) {
		return 1
	} else {
		return -1
	}
}

func openWriteFile (sOpenFileName string) os.File {		

	file, err := os.OpenFile(
				sOpenFileName, 
				os.O_CREATE|os.O_RDONLY|os.O_APPEND,
				os.FileMode(0644))
		if err != nil {
			fmt.Println(err)
		}
		
		return *file
}

