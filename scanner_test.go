package main


import (
	"fmt"
	"testing"
	"os"
	"reflect"
	"net"
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

func TestConnect(t *testing.T) {
	// test port retrieval
	TestPorts, err := GetAddrs(_gethostname())
	if err != nil {
		t.Error("Error getting ports: ", err)
	}

	connection := connect("tcp", TestPorts[0])
	var conn net.Conn
	
	// test return type
	if reflect.TypeOf(connection) != reflect.TypeOf(conn) {
		t.Error("expected type: []string \n result: ", reflect.TypeOf(connection))
	}
}