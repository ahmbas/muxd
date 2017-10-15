package mux

import (
	"fmt"
	"os"

	"github.com/ahmbas/muxd/protocols"
	"github.com/ahmbas/muxd/protocols/redis"
	"github.com/ahmbas/muxd/protocols/tcp"
	"github.com/ahmbas/muxd/protocols/udp"
)

//Mux (Input, Output)
type Mux struct {
	Input  protocols.BaseConnection
	Output protocols.BaseConnection
}

//Run forever
func (mux Mux) Run() {
	c := make(chan []byte)

	for b := range mux.Input.Rx(c) {
		mux.Output.Tx(b)
	}

}

//Opts (all opts)
type Opts struct {
	Protocol string
	Host     string
	Port     int
	Channel  string
}

func GetConnection(o Opts) protocols.BaseConnection {

	switch o.Protocol {
	case "redis":
		return redis.Connection{
			Host:    o.Host,
			Port:    o.Port,
			Channel: o.Channel,
		}
	case "tcp", "tcp4", "tcp6":
		return tcp.Connection{
			Protocol: o.Protocol,
			Host:     o.Host,
			Port:     o.Port,
		}
	case "udp", "udp4", "udp6":
		return udp.Connection{
			Protocol: o.Protocol,
			Host:     o.Host,
			Port:     o.Port,
		}
	default:
		fmt.Printf("Invalid args. Run with --help")
		os.Exit(1)
	}
	return nil

}
