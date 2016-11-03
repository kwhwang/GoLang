package main

import (
	"fmt"
	"os"
	//"io"
	"unsafe"
	"bytes"
	"encoding/binary"
)
type  st_exit_status struct {
	e_termination  uint16  		
	e_exit uint16         		
}

type st_wtmp struct {
	ut_type uint32
	ut_pid  uint32			
	ut_line [32] byte 
	ut_id   [4] byte              	
	ut_user [32] byte  
	ut_host [256] byte  	
	ut_exit st_exit_status 		
	ut_session uint32        	
	st_ut_tv struct {
		tv_sec  uint32          
		tv_usec uint32         
	}
	ut_addr_v6 [4] uint32       
	unused [20] byte           
}

func main() {
	sFileName, sTail := checkArgu ()

	fmt.Printf("FILENAME = [%s]\n", sFileName)
	fmt.Printf("OPTION   = [%s]\n", sTail)


	//file := openFile(sFileName)

	file, err := os.Open(sFileName)
	
	wtmp := st_wtmp{}

	iSize := unsafe.Sizeof(wtmp)

	fmt.Printf("size=[%d]\n", iSize)

	data := readNextBytes(file,iSize)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.LittleEndian, &wtmp)

	fmt.Printf("",err)

	fmt.Printf("ut_type:[%d], ut_pid:[%d], ut_line:[%s], ut_id:[%s], ut_user:[%s], ut_host:[%s], e_termination:[%d], e_exit:[%d], ut_session:[%d], tv_sec:[%d], tv_usec:[%d], ut_addr_v6:[%d], unused:[%s]\n", 
				wtmp.ut_type, 
				wtmp.ut_pid, 
				wtmp.ut_line[:], 
				wtmp.ut_id[:], 
				wtmp.ut_user[:],
				wtmp.ut_host[:], 
				wtmp.ut_exit.e_termination, 
				wtmp.ut_exit.e_exit, 
				wtmp.ut_session, 
				wtmp.st_ut_tv.tv_sec, 
				wtmp.st_ut_tv.tv_usec, 
				wtmp.ut_addr_v6,
				wtmp.unused[:])

	/*
	if err != nil {
		if err == io.EOF {
			return			
		}
	} else {
		fmt.Printf("ut_type:[%d], ut_pid:[%d], ut_line:[%s], ut_id:[%s], ut_user:[%s], ut_host:[%s], e_termination:[%d], e_exit:[%d], ut_session:[%d], tv_sec:[%d], tv_usec:[%d], ut_addr_v6:[%d], unused:[%s]\n", 
				wtmp.ut_type, 
				wtmp.ut_pid, 
				wtmp.ut_line[:], 
				wtmp.ut_id[:], 
				wtmp.ut_user[:],
				wtmp.ut_host[:], 
				wtmp.ut_exit.e_termination, 
				wtmp.ut_exit.e_exit, 
				wtmp.ut_session, 
				wtmp.st_ut_tv.tv_sec, 
				wtmp.st_ut_tv.tv_usec, 
				wtmp.ut_addr_v6,
				wtmp.unused[:])
	}
	*/
	
}

func checkArgu () (string,string) {

	var sFileName string
	var sTail string

	if ( len(os.Args) == 1 ) {				
		sFileName = fmt.Sprintf("%s", "/var/log/wtmp")
	} else if ( len(os.Args) == 2 ) {		
		if ( os.Args[1] == "-t" ) {
			sTail = fmt.Sprintf("%s", os.Args[1])			
			sFileName = fmt.Sprintf("%s", "/var/log/wtmp")
		} else {
			sTail = ""			
			sFileName = fmt.Sprintf("%s", os.Args[1])
		}
	} else if ( len(os.Args) == 3 ) {
		if ( os.Args[1] == "-t" ) {
			sTail = fmt.Sprintf("%s", os.Args[1])
			sFileName = fmt.Sprintf("%s", os.Args[2])
		} else {
			sTail = fmt.Sprintf("%s", os.Args[2])
			sFileName = fmt.Sprintf("%s", os.Args[1])
		}				
	} else {
		os.Exit(2)
	}

	return sFileName, sTail
}
/*
func openFile (sOpenFileName string) *os.File {		

	file, err := os.Open(sOpenFileName)
		if err != nil {
			fmt.Println(err)
		}
		
		return file
}
*/

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		//log.Fatal(err)
	}

	return bytes
}
