package salvo

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
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

// Ulimit gets the hard ceiling # of pthreads
func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

// GetRange gets a range
func GetRange(min int, max int) []int {
	var result []int
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return result
}

// PortWorker is one scan process
func PortWorker(protocol string, tgt string) (string, bool) {
	fmt.Println("Scanning port " + tgt)
	// randomized timeout
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)

	// attempt to dial
	connection, err := net.DialTimeout(protocol, tgt, time.Duration(10))
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(time.Duration(10))
			PortWorker(protocol, tgt)
		}
		return tgt, false
	}

	if connection != nil {
		// connection succeeded
		return tgt, true
	}
	defer connection.Close()
	return tgt, false
}

// PortScan dials host:port addresses and returns a list of ports
func PortScan(protocol string, tgthost string, portMin int, portMax int) (result map[string]bool) {

	// port range
	ports := GetRange(portMin, portMax)
	// randomize port access
	ports = _shuffleOrder(ports)

	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(Ulimit() - 400)
	resultchan := make(chan status)

	// parallel for loop pattern
	for _, v := range ports {
		wg.Add(1)
		sm.Acquire(context.TODO(), 1)
		address := net.JoinHostPort(tgthost, string(v))
		// go routine for each worker
		go func(protocol string, address string) {
			defer wg.Done()
			defer sm.Release(1)
			addr, open := PortWorker(protocol, address)
			result[addr] = open
		}(protocol, address)
	}
	// wait for goroutines to end and return
	wg.Wait()
	return result
}
