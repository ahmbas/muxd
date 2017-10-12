package udp

import (
	"fmt"
	"net"
	"os"
)

//Rx receive tcp traffic
func (input Connection) Rx(c chan []byte) chan []byte {

	go func() {
		addr, err := net.ResolveUDPAddr(
			input.Protocol,
			fmt.Sprintf("%v:%v", input.Host, input.Port),
		)
		if err != nil {
			fmt.Println("Couldnt resolve:", err.Error())
		}
		conn, err := net.ListenUDP(input.Protocol, addr)
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		for {
			buff := make([]byte, 65200)
			n, _, err := conn.ReadFromUDP(buff)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Printf("Got %v", string(buff[:n]))
			c <- buff[:n]
		}

	}()
	return c
}
