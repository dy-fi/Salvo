package main

import (
	"os"
	"fmt"

)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: can't get host \n%v", err)
	}

	conns := make(map[string]string)
	addrs,_ := GetAddrs(name)
	ports := []int{80, 443}

	for _,v := range addrs {
		conns = ConnScan("tcp", v, ports) 
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for _,v := range conns {
			fmt.Println(v) 
		}
	}
}

