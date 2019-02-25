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

	conns := []string{}
	ports := []int{}

	for i := 8000; i > 0; i-- {
		ports = append(ports, i)
	}

	for _, v := range addrs {
		conns = PortScan("tcp", v, ports)
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for _, v := range conns {
			fmt.Println(v + "\n")
		}
	}
}
