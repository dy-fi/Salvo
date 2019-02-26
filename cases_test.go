package salvo

import (
	"os"
	"net"
)

var host,_ = os.Hostname()
var addrs, err = net.LookupHost(host)
var strlist []string

var PortScanCases = []struct {
	description string
	protocol string
	ports 	 []int
}{
	{
		"Basic internet ports",
		"tcp",
		_getlist(2000),
	},
	{

	},
}