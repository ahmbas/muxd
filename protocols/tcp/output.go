package tcp

import (
	"net"

	log "github.com/sirupsen/logrus"
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
		log.Warnf(
			"Error Connecting to %v %v",
			output,
			err.Error(),
		)
		return
	}

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

	conn.Close()

}
