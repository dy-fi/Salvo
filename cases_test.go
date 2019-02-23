package main

import (
	"os"
	"net"
)

var host,_ = os.Hostname()
var currHost, err = net.LookupHost(host)
var strlist []string

var GetAddrsCases = []struct {
	input string
	output interface{}
	error bool 
}{
	{
		host,
		currHost,
		false,
	},
	{
		"asdf",
		nil, 
		true,
	},
}