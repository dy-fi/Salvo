package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"github.com/dy-fi/salvo/salvo"
)

func main() {
	start := time.Now()
	// environment data
	host := os.Args[1]

	fmt.Printf("Scanning on host %v...\n\n", host)

	// resolve args
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println("Error: can't get host")
		panic(err)
	}

	portMin, err := strconv.Atoi(os.Args[2])
	portMax, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error: cant resolve ports: %v", err)
		panic(err)
	}

	ls := salvo.PortScan("tcp", addrs[0], portMin, portMax)
	fmt.Printf("\n%v ports scanned in %v\n\n", portMax-portMin, time.Since(start))
	fmt.Println("OPEN PORTS")
	fmt.Println(ls)
}

