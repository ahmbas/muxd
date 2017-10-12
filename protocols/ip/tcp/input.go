package tcp

import (
	"fmt"
	"net"
	"os"
)

//Rx receive tcp traffic
func (input Connection) Rx(c chan []byte) chan []byte {

	go func() {
		l, err := net.Listen(input.Protocol, fmt.Sprintf("%v:%v", input.Host, input.Port))
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			buff := make([]byte, 65200)
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Printf("Got %v", string(buff[:n]))
			c <- buff[:n]
		}

	}()
	return c
}
