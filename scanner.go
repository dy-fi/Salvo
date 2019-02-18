package main

import (
	"net"
	"sync"
	"errors"
)

type connection struct {
	conn *net.Conn
	err *error 
}


// GetAddrs returns a slice of addresses on the given port
func GetAddrs(hostname string) []string {

	var ports []string
	ports,err := net.LookupHost(hostname)

	if err != nil {
		errors.New("Couldn't resolve host")
	}
	return ports 
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

		go func(v string) {
			defer wg.Done()
			connection := connect("tcp", address)
			if connection != nil {
				results = append(results, address)
			}
		}(v);
	}
	wg.Wait()

	return results
}
