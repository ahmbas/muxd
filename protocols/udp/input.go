package udp

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

//Rx receive udp traffic
func (input Connection) Rx(c chan []byte) chan []byte {

	log.Debugf("Starting Input on %v", input)

	go func() {
		addr, err := net.ResolveUDPAddr(
			input.Protocol,
			fmt.Sprintf("%v:%v", input.Host, input.Port),
		)
		if err != nil {
			log.Warnf("Couldnt resolve %v %v", input, err.Error())
		}
		conn, err := net.ListenUDP(input.Protocol, addr)
		if err != nil {
			log.Fatalf("Error accepting %v %v ", input, err.Error())
			return
		}
		for {
			buff := make([]byte, 65200)
			n, _, err := conn.ReadFromUDP(buff)
			if err != nil {
				log.Warnf("Error reading %v %v", input, err.Error())
			}
			log.Debugf("Got %v", string(buff[:n]))
			c <- buff[:n]
		}

	}()
	return c
}
