package salvo

import (
	"fmt"
	"os"
	"net"
)

func main() {
	// environment data
	host, err := os.Hostname()
	if err != nil {
		return
	}
	addrs,err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("Error: can't get host \n%v", err)
		return
	}

	conns := make(map[string]*net.Conn)
	ports := _getlist(8000)

	for _, v := range addrs {
		conns = PortScan("tcp", v, ports)
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for k := range conns {
			fmt.Println(k + "\n")
		}
	}
}
