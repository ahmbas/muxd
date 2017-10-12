package tcp

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

//Rx receive tcp traffic
func (input Connection) Rx(c chan []byte) chan []byte {

	log.Debugf("Starting Input on %v", input)

	go func() {
		l, err := net.Listen(input.Protocol, fmt.Sprintf("%v:%v", input.Host, input.Port))
		if err != nil {
			log.Fatalf("Error listening on %v %v", input, err.Error())
			os.Exit(1)
		}
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatalf("Error accepting %v %v ", input, err.Error())
				os.Exit(1)
			}
			buff := make([]byte, 65200)
			n, err := conn.Read(buff)
			if err != nil {
				log.Warnf("Error accepting %v %v ", input, err.Error())
			}
			log.Debugf("Got %v", string(buff[:n]))
			c <- buff[:n]
		}

	}()
	return c
}
