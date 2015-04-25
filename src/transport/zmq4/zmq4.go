package zmq4

import (
	zmq "github.com/pebbe/zmq4"
)

type Transport struct {
	Context zmq.Context
    Socket zmq.Socket
}

func NewTransport() (t Transport) {
	t.Context = zmq.NewContext()
	t.Socket = context.NewSocket(zmq.SUB)
	t.Socket.Bind("epgm://eth0;239.192.1.1:55555")
}

func (t Transport) Run() {
	for {
		msg, _ := t.Socket.Recv(0)
		println("Received ", string(msg))
	}
}
