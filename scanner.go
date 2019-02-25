package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"math/rand"
	"time"
)

// shuffles port access order to obfuscate 
func _shuffleOrder(src []string) []string {
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
    	dest[v] = src[i]
	}
	return dest
}

// PortWorker is one scan process
func PortWorker(protocol string, tgt string, verb ...bool) {
	// randomized timeout
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)

	// attempt to dial
	connection, err := net.Dial("tcp", tgt)
	status, _ := bufio.NewReader(connection).ReadString('\n')
	defer connection.Close()

	if verb[1] {
		if connection != nil && err == nil {
			// connection succeeded
			fmt.Println("√ ", status)
		} else {
			// connection failed
			fmt.Println("X", status)
		}
	}
}

// PortScan dials host:port addresses and returns a list of successes
func PortScan(protocol string, tgthost string, tgtports []int, verb ...bool) (result map[string]string) {
	var wg sync.WaitGroup
	verbose := false

	if verb != nil {
		if verb[0] {
			verbose = true
		}
	}

	// parallel for loop pattern
	for _, v := range tgtports {
		wg.Add(1)
		address := net.JoinHostPort(tgthost, string(v))

		go func(address string) {
			defer wg.Done()
			PortWorker(protocol, address, verbose)
		}(address)
	}
	// wait for goroutines to end and return
	wg.Wait()
	return
}

