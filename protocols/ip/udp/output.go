package udp

import (
	"fmt"
	"net"
)

//Tx send UDP traffic
func (output Connection) Tx(msg []byte) {

	addr := net.UDPAddr{
		IP:   net.ParseIP(output.Host),
		Port: output.Port,
	}

	conn, err := net.DialUDP(
		output.Protocol,
		nil,
		&addr,
	)

	_, err = conn.Write(msg)
	if err != nil {
		fmt.Printf("Could not forward the following data %v", string(msg))
	}

}
