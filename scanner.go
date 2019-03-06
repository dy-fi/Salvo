package salvo

import (
	"math/rand"
	"net"
	"sync"
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
func PortWorker(protocol string, tgt string) (string, bool) {
	// randomized timeout
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)

	// attempt to dial
	connection, err := net.Dial(protocol, tgt)
	if err != nil {
		return tgt, false
	}

	if connection != nil {
		// connection succeeded
		return tgt, true
	}
	defer connection.Close()
	return tgt, false

}

// PortScan dials host:port addresses and returns a list of successes
func PortScan(protocol string, tgthost string, tgtports []int) (result map[string]bool) {
	var wg sync.WaitGroup

	// randomize port access
	ports := _shuffleOrder(tgtports)

	// parallel for loop pattern
	for _, v := range ports {
		wg.Add(1)
		address := net.JoinHostPort(tgthost, string(v))

		go func(address string) {
			defer wg.Done()
			address, status := PortWorker(protocol, address)
			result[address] = status
		}(address)
	}
	// wait for goroutines to end and return
	wg.Wait()
	return result
}