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

	ports, err := GetAddrs(name)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for _,v := range ConnScan(name, ports) {
		fmt.Println(v) 
	}
}

