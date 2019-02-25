package main

import (
	"net"
	"sync"
	"math/rand"
	"time"
)

// shuffles a list of integers
func _shuffleOrder(src []int) []int {
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
    	dest[v] = src[i]
	}
	return dest
}

// PortWorker is one scan process
func PortWorker(protocol string, tgt string) (bool, string) {
	// randomized timeout
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)

	// attempt to dial
	connection, err := net.Dial("tcp", tgt)
	if err != nil {
		return false, ""
	}
	
	if connection != nil {
		// connection succeeded
		return true, tgt
	}
	defer connection.Close()
	return false, ""
	
}

// PortScan dials host:port addresses and returns a list of successes
func PortScan(protocol string, tgthost string, tgtports []int) (result []string) {
	var wg sync.WaitGroup
	

	// randomize port access
	ports := _shuffleOrder(tgtports)

	// parallel for loop pattern
	for _, v := range ports {
		wg.Add(1)
		address := net.JoinHostPort(tgthost, string(v))

		go func(address string) {
			defer wg.Done()
			status, address := PortWorker(protocol, address)
			if status {
				result = append(result, address)
			}
		}(address)
	}
	// wait for goroutines to end and return
	wg.Wait()
	return 
}

