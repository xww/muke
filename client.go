package main

import (
	"flag"
	"net"
	"strconv"

	"bytes"
	"encoding/binary"
)

var chost = flag.String("chost", "0.0.0.0", "chost")
var cport = flag.String("cport", "5555", "cport")

func main() {
	flag.Parse()
	conn, _ := net.Dial("tcp", *chost + ":" + *cport)
	tcpConn, _ := conn.(*net.TCPConn)
	tcpConn.SetNoDelay(true)

	for i := 0; i < 500; i++ {
		content := "set:" + strconv.Itoa(i) + ":" + strconv.Itoa(i)

		lenByte := make([]byte, 8)
		binary.BigEndian.PutUint64(lenByte, uint64(len(content)))
		var buffer bytes.Buffer
		buffer.Write(lenByte)
		buffer.Write([]byte(content))
		conn.Write(buffer.Bytes())

	}

}
