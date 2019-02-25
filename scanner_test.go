package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func _gethostname() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Can't resolve any hostnames (you might need to connect to the internet)")
	}
	return name
}

func TestPortScan(t *testing.T) {

	conn := PortScan("tcp", _gethostname(), []int{80, 443})

	// test valid output
	if len(conn) < 1 {
		t.Fatal("Could not detect any ports when at least some should be accessible")
	}
	// test cases
	for _, tc := range GetAddrsCases {
		output := PortScan("tcp", host, []int{80, 443})
		if tc.error {
			var _ error = err

			if err == nil {
				t.Fatalf("GetAddrs(%q)Should have thrown error- \ngot: %d \nwant: %d", tc.input, output, tc.output)
			}
		} else {
			if reflect.TypeOf(output) != reflect.TypeOf(tc.output) {
				t.Fatalf("GetAddrs(%q) unexpected output- \ngot: %d \nwant: %d", tc.input, output, tc.output)
			}
			if err != nil {
				t.Fatalf("GetAddrs(%q) unexpected error- %d", tc.input, err)
			}
		}
	}
}

func BenchmarkPortScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestPorts := PortScan("tcp", _gethostname(), []int{80, 443})
		if TestPorts == nil {
			fmt.Println("Couldn't find a host to benchmark on")
			return
		}
	}
}
