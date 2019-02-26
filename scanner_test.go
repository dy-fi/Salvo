package salvo

import (
	"fmt"
	"os"
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

	conn := PortScan("tcp", _gethostname(), _getlist(4000))

	// test valid output
	if len(conn) < 1 {
		t.Fatal("Could not detect any ports when at least some should be accessible")
	}
	// test cases
	for _, tc := range PortScanCases {
		output := PortScan(tc.protocol, host, tc.ports)

		if output == nil {
			t.Fatal("Expected output but none was given")
		}
	}
}

func BenchmarkPortScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestPorts := PortScan("tcp", _gethostname(), _getlist(4000))
		if TestPorts == nil {
			fmt.Println("Couldn't find a host to benchmark on")
			return
		}
	}
}
