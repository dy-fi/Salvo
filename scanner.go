package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"sync"
)

// GetAddrs returns a slice of addresses on the given host
func GetAddrs(hostname string) ([]string, error) {

	var addresses []string
	addresses, err := net.LookupHost(hostname)

	if err != nil {
		return nil, errors.New("Couldn't resolve host")
	}
	return addresses, nil
}

// ConnScan dials host:port addresses and returns a list of successes
func ConnScan(protocol string, tgthost string, tgtports []int, verb ...bool) (result map[string]string) {
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
