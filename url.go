package main
 
import (
    "fmt"
    //"io/ioutil"
    "net/http"
)

func main() {	
   
    for i := 30; i<50; i++ {
        
        for j := 0; j<100; j++ {

        var tmp string

	uuid := "F89DC6A27CA329AB37565114AB7D7EA4igloosec"

	tmp = fmt.Sprintf("http://192.168.150.81:8080/spider-x/AgentInfoSetup?kind=save&proxy_info=16250,-,-,1&parent_id=16249&parent_name=Test&mgr_ip=192.168.150.82&product=Other&type=V&proxy_type=S&user_id=igloosec&name=VIRTUAL_192.100.%d.%d&agent_ip=192.100.%d.%d&uuid=%s", i,j,i,j, uuid)

	fmt.Printf("%s\n", tmp);
   		
        http.Get(tmp)

        }
    }
   
}
