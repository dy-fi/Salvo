package main

import (
	"os"
	"fmt"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	addrs, err := GetAddrs(name)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	ports := ConnScan(name, addrs) 

	if len(ports) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for _,v := range ports {
			fmt.Println(v) 
		}
	}
}

