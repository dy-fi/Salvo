package salvo

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// environment data
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting host: %v", err)
	}
	addrs, err  := net.LookupHost(host)
	if err != nil {
		fmt.Printf("Error looking up host: %v", err)
	}
	fmt.Println(string(len(addrs)))
	var conns []string
	ports := _getlist(8000)

	for _, v := range addrs {
		conns = PortScan("tcp", v, ports)
	}
	
	for k := range conns {
		fmt.Printf("%v\n", k)
	}
	
}
