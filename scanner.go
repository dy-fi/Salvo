package main

import (
	"net"
	"sync"
	"errors"
)


// GetAddrs returns a slice of addresses on the given port
func GetAddrs(hostname string) ([]string, error) {

	var ports []string
	ports,err := net.LookupHost(hostname)

	if err != nil {
		return nil, errors.New("Couldn't resolve host")
	}
	return ports, nil
}

// connect attempts a connection
func connect(network string, address string) (*net.Conn) {
	conn,err := net.Dial(network, address)
	
	if err != nil {
		return nil
	}
	return &conn 
}

// ConnScan dials host:port addresses on a host and list of successes
func ConnScan(tgthost string, tgtports []string) ([]string){

	var results []string 
	var wg sync.WaitGroup
	
	// parallel for loop pattern
	for _,v := range tgtports {
		wg.Add(1)
		address := net.JoinHostPort(tgthost, v)

		go func(address string) {
			defer wg.Done()
			connection := connect("tcp", address)
			if connection != nil {
				results = append(results, address)
			}
		}(address)
	}
	wg.Wait()

	return results
}
