package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

// PortScan dials host:port addresses and returns a list of successes
func PortScan(protocol string, tgthost string, tgtports []int, verb ...bool) (result map[string]string) {
	var wg sync.WaitGroup

	// parallel for loop pattern
	for _, v := range tgtports {
		wg.Add(1)
		address := net.JoinHostPort(tgthost, string(v))

		go func(address string) {
			defer wg.Done()

			// attempt to dial
			connection, err := net.Dial("tcp", address)
			status, _ := bufio.NewReader(connection).ReadString('\n')
			defer connection.Close()

			// success
			if connection != nil && err == nil {
				result[address] = status

				// excluded without flag for optimization
				if verb[1] {
					fmt.Println("√ ", status)
				}
				// fail
			} else {
				if verb[1] {
					fmt.Println("X", status)
				}
			}
		}(address)
	}
	// wait for goroutines to end and return
	wg.Wait()
	return
}

// FullScan is simply 
