package tcp

import (
	"fmt"
	"net"
	"os"
)

//Tx send tcp traffic
func (output Connection) Tx(msg []byte) {

	addr := net.TCPAddr{
		IP:   net.ParseIP(output.Host),
		Port: output.Port,
	}
	conn, err := net.DialTCP(
		output.Protocol,
		nil,
		&addr,
	)
	if err != nil {
		fmt.Printf("Error Connecting to %v:%v %v", output.Host, output.Port, err)
		os.Exit(1)
	}
	conn.Write(msg)
	conn.Close()

}
