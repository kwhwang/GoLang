package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"
)
/*		FUTURE USE
type SPCOLLECT_INFO struct {
	sAddr string
	sTime string
	sLastRecvTime string
	sReplaceCommand string
	iCount	int
	iFd	int
}
*/

type SPCOLLECT_CONF struct {
	sPort		string
	sFileName	string
	//iReadCount	int			//FUTURE USE
	//iSleepTime	int			//FUTURE USE
}

var gpspcollect_info = new(SPCOLLECT_CONF)


const UDP_SOCKET_BUFFER_SIZE = 83886080
const RECV_BUFFER_SIZE = 10240



func main() {
	var sPort string 
	var sDataBuf string
	var sNowTime string
	var sWriteTime string
	var i int
	var iReturn int
	var sOldTime string
	var fFile os.File
	
	
	sPort = SetDefaultValue()
	
	ServerAddr,err := net.ResolveUDPAddr("udp", sPort)						// UDP 소켓 생성
	ServerConn, err := net.ListenUDP("udp", ServerAddr)						// Listen
	CheckError(err)
	defer ServerConn.Close()

	err = ServerConn.SetReadBuffer(UDP_SOCKET_BUFFER_SIZE)								// UDP Buffer 사이즈 증가

	sRecvBuf := make([]byte, RECV_BUFFER_SIZE)											// 버퍼 생성
	
	sOldTime = convertStringtoTime(gpspcollect_info.sFileName)				// 프로그램 시작 시간 및 현재시간 파일 생성 (초기화)
	fFile = openWriteFile(sOldTime)

	for {
		sNowTime = convertStringtoTime(gpspcollect_info.sFileName)			// 파일명에 yyyymmddhhmi 값으로 현재 시간을 체크

 		n,addr,err := ServerConn.ReadFromUDP(sRecvBuf)						// udp 생성

		if err != nil {														// 에러 체크 
			fmt.Println("Error: ", err)
		}
		
		iReturn = compareOldNowTime(sOldTime, sNowTime) 					//Write중인 파일 명의 갱신이 필요한지 확인 (분단위가 바뀌었는지 확인) 

		if ( iReturn != 1 ) {												// 1이 아닌 경우 시간 변경. 그외에는 같은 분대 00초~59초, 
			fmt.Printf("File Switching. Old=[%s], Now=[%s]\n", sOldTime, sNowTime)
			defer fFile.Close()												
			fFile = openWriteFile(sNowTime)
			sOldTime = fmt.Sprintf("%s", sNowTime)
		} 
		
		sWriteTime  = convertStringtoTime("yyyymmdd hhmiss")				// 파일에 저장될 현재시간 (yyyymmddhhmiss)

		sDataBuf = fmt.Sprintf("[%s][%s][%s]\n",addr.IP, sWriteTime, string(sRecvBuf[0:n]))		// 파일 내용 합치기
		
		n, tmp := fFile.Write([]byte(sDataBuf))													//파일 쓰기

		i++
		if ( i > 100 ) {																		// flush 카운트. (Timer를 두고 주기적으로 flush 하는 방안이 더 좋음)
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

	//gpspcollect_info.iReadCount		= 5000								//FUTURE USE
	//gpspcollect_info.iSleepTime		= 1000	
	//fmt.Println(gpspcollect_info.iReadCount)
	//fmt.Println(gpspcollect_info.iSleepTime)

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

func compareOldNowTime(sOldTime string, sNowTime string) int {		//현재 시간과 이전 파일 시간 비교
	
	if ( strings.Contains(sNowTime, sOldTime) ) {
		return 1
	} else {
		return -1
	}
}

func openWriteFile (sOpenFileName string) os.File {		//파일 오픈 함수

	file, err := os.OpenFile(
				sOpenFileName, 
				os.O_CREATE|os.O_RDONLY|os.O_APPEND,
				os.FileMode(0644))
		if err != nil {
			fmt.Println(err)
		}
		
		return *file
}

