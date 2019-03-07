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
		return
	}

	fmt.Println(host)

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("Error: can't get host \n%v", err)
	}

	var conns map[string]bool

	for _, v := range addrs {
		conns = PortScan("tcp", v, 1000, 3000)
	}

	for i, v := range conns {
		if v {
			fmt.Println(i + ": success \n")
		} else {
			fmt.Println(i + ": failure \n")
		}
	}

}
