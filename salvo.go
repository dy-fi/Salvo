package main

import (
	"fmt"
	"github.com/dy-fi/salvo"
	"net"
	"os"
)

func _getlist(n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

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
		return
	}
	fmt.Print(string(len(addrs)))

	var conns map[string]bool
	ports := _getlist(8000)

	for _, v := range addrs {
		conns = salvo.PortScan("tcp", v, ports)
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for i, v := range conns {
			if v {
				fmt.Println(i + ": success \n")
			} else {
				fmt.Println(i + ": failure \n")
			}
			
		}
	}
}
