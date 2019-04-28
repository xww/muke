package demo

import (
	"flag"
	"net"
	"fmt"

	//"encoding/binary"
	"strings"

)

var host = flag.String("host", "0.0.0.0", "host")
var port = flag.String("port", "5555", "port")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *host + ":" + *port)
	if err != nil{
		fmt.Println(err)
	}

	defer l.Close()
	fmt.Println("listen on " + *host + ":" + *port)
	for {
		conn, _ := l.Accept()

		fmt.Printf("receive message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {

	kvmap := make(map[string]string)

	defer func() {
		conn.Close()

	}()

	for {
		//buf := make([]byte, 8)
		//cnt, _ := conn.Read(buf)
		//contentLen := uint64(binary.BigEndian.Uint64(buf))
		//buf = make([]byte, contentLen)
		//cnt, _ = conn.Read(buf)

		buf := make([]byte, 4096)
		cnt, _ := conn.Read(buf)

		if cnt == 0{
			break
		}

		fmt.Println("receive:" + string(buf))
		instr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(instr, ":")
		switch inputs[0] {
		case "set":
			kvmap[inputs[1]] = inputs[2]
		default:
			fmt.Printf("unsupport command:%s\n", inputs[0])

		}
	}
	fmt.Println(kvmap)

}
