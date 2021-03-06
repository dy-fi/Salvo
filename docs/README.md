# Salvo
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/597d7ec3d7da451682ba7d633312efad)](https://www.codacy.com/app/dy-fi/Salvo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=dy-fi/Salvo&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/dy-fi/Salvo)](https://goreportcard.com/report/github.com/dy-fi/Salvo)

A package leverages Go concurrency for more efficient scanning.  

* randomized request times
* randomized port order

Coming Soon:
* Source IP spoofing by manual packet construction
* file dumping
___ 

## Doc Site
Hosted on its own site [here](https://dy-fi.github.io/Salvo/#/)

### Author
Dylan Finn - [LinkedIn](https://www.linkedin.com/in/dylan-finn-a36b9614b/) - [GitHub](https://github.com/dy-fi)

### License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

___

### Installation Instructions
`go get github.com/dy-fi/salvo`

### Quick Start Scan Recipe

This is a TCP full scan of the current host on port 80 (http) and 443 (https)

```go
import (
	"fmt"
	"os"
	"net"
)

func main() {
	// environment data
	host, err := os.Hostname()
	addrs,err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("Error: can't get host \n%v", err)
	}

	conns := []string{}
	ports := []int{}

	for i := 8000; i > 0; i-- {
		ports = append(ports, i)
	}

	for _, v := range addrs {
		conns = PortScan("tcp", v, ports)
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for _, v := range conns {
			fmt.Println(v + "\n")
		}
	}
}

```

### Documentation

#### PortScan

`func PortScan(protocol string, tgthost string, tgtports []int) ([]string)`

PortScan dials host:port addresses concurrently and returns a list of successes 

| param             | description            |
|-------------------|:----------------------:|
| `protocol string` |     i.e. "tcp"         |
| `tgthost string`  | target host hostname   |
| `tgtports`        | ports to scan          |

