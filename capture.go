package salvo

import (
	"io/ioutil"
	"net"
	"time"
)

// Capture struct to operate on a captured connection
type Capture struct {
	// Stream captured connection
	Stream net.Conn
	// Active status bool to hold capture state
	Active bool
	// Receiver channel to read from
	Receiver chan (string)
}

// NewCapture acts as the capture constructor
// 0 value for time means read/write will not time out
func NewCapture(conn net.Conn, timeOption ...time.Time) Capture {
	var tgt Capture
	var receiver chan (string)
	tgt.Stream = conn
	tgt.Receiver = receiver
	tgt.Stream.SetDeadline(timeOption[0])

	return tgt
}

// Read reads the connection and passes it to the Receiver channel
func (c *Capture) Read() {
	b, _ := ioutil.ReadAll(c.Stream)
	c.Receiver <- string(b)
}

// IPListen checks to see if a connection is IP type and then
func (c *Capture) IPListen(network string) {

}
