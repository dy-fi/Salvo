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

func TestGetAddrs(t *testing.T) {

	// test port retrieval
	TestPorts, err := GetAddrs(_gethostname())
	if err != nil {
		t.Error("Error getting ports: ", err)
	}

	var s []string
	// test return type
	if reflect.TypeOf(TestPorts) != reflect.TypeOf(s) {
		t.Fatal("expected type: []string \n result: ", reflect.TypeOf(TestPorts))
	}
	// test valid output
	if len(TestPorts) < 1 {
		t.Fatal("Could not detect any ports when at least some should be accessible")
	}

	for _, tc := range GetAddrsCases {
		output, err := GetAddrs(tc.input)
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

func TestConnScan(t *testing.T) {

}

func BenchmarkGetAddrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestPorts, _ := GetAddrs(_gethostname())
		if TestPorts == nil {
			fmt.Println("Couldn't find a host to benchmark on")
			return
		}
	}
}
