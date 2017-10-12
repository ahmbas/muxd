package muxd

import "github.com/ahmbas/muxd/protocols"

//Mux
type Mux struct {
	input  protocols.BaseConnection
	output protocols.BaseConnection
}

func (mux Mux) Run() {
	c := make(chan []byte)

	for b := range mux.input.Rx(c) {
		mux.output.Tx(b)
	}

}
