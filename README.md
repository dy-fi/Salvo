# Salvo
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/597d7ec3d7da451682ba7d633312efad)](https://www.codacy.com/app/dy-fi/Salvo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=dy-fi/Salvo&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/dy-fi/Salvo)](https://goreportcard.com/report/github.com/dy-fi/Salvo)

A package leverages Go concurrency for more efficient scanning.  

___ 

### Documentation
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
    "os"
    "fmt"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: can't get host \n%v", err)
	}

	conns := make(map[string]string)
	addrs,_ := GetAddrs(name)
	ports := []int{80, 443}

	for _,v := range addrs {
		conns = ConnScan("tcp", v, ports) 
	}

	if len(conns) == 0 {
		fmt.Println("No ports were detected")
	} else {
		for _,v := range conns {
			fmt.Println(v) 
		}
	}
}
```

___

### Documentation


#### ConnScan

`func ConnScan(protocol string, tgthost string, tgtports []int, verb ...bool) (map[string]string)`

ConnScan dials host:port addresses and returns a list of successes 

| param             | description            |
|-------------------|:----------------------:|
| `protocol string` |     i.e. "tcp"         |
| `tgthost string`  | target host hostname   |
| `tgtports`        | ports to scan          |
| `verb ...bool`    | verbose option         |


#### GetAddrs

`func GetAddrs(hostname string) ([]string, error)`

GetAddrs returns a slice of addresses on the given host

| param             | description            |
|-------------------|:----------------------:|
| `hostname string` | address or resolvable string         |

