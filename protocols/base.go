package protocols

// BaseConnection base
type BaseConnection interface {
	Tx(msg []byte)
	Rx(c chan []byte) chan []byte
}
