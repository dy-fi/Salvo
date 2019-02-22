package main


import (
	"fmt"
	"testing"
	"os"
	"reflect"
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
	if reflect.TypeOf(TestPorts) !=  reflect.TypeOf(s) {
		t.Error("expected type: []string \n result: ", reflect.TypeOf(TestPorts))
	}
	// test valid output
	if len(TestPorts) < 1 {
		t.Error("Could not detect any ports when at least some should be accessible")
	}
}

func BenchmarkGetAddrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestPorts,_ := GetAddrs(_gethostname())
		if TestPorts == nil {
			fmt.Println("Couldn't find a host to benchmark on")
			return
		}
	}
}
