package amqp09

import (
	amqp "github.com/streadway/amqp"
	"os"
	"protocol"
	"github.com/op/go-logging"
	"time"
)

var log = logging.MustGetLogger("example")

type Transport struct {
	Conn *amqp.Connection
}

func NewTransport()  Transport {
	//t.Conn.err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn,err := amqp.Dial(os.Getenv("AMQP_URL"))
	var t Transport
	t.Conn = conn

	if err != nil {
		log.Fatal("connection.open: %s", err)
	}
//	err = c.ExchangeDeclare("logs", "topic", true, false, false, false, nil)
//	if err != nil {
//		log.Fatalf("exchange.declare: %s", err)
	//	}
	return t
}

//func (t Transport) Run() {
//	firehose, err := t.Conn.Consume("firehose", "", true, false, false, false, nil)
//	if err != nil {
//		log.Fatal("basic.consume: %v", err)
//	}
//}

func (t *Transport) SendEvent(ev protocol.Event, endpoint string, path string) {
	_ = ev
	c,_ := t.Conn.Channel()
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte("Go Go AMQP!"),
	}
	c.Publish(endpoint, path, false, false, msg)
}
