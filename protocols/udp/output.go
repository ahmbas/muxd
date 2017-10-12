package udp

import (
	"net"

	log "github.com/sirupsen/logrus"
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
		log.Warnf(
			"Could not forward the following data %v %v %v",
			output,
			string(msg),
			err.Error(),
		)
		return
	}

	log.Debugf("Sent %v to %v", string(msg), output)

}
