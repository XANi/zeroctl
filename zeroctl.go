package main

import (
	//	"fmt"
	"crypto/rand"
	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"net"
	"os"
	"protocol"
	"strings"
	"time"
	amqp09 "transport/amqp09"
)

var log = logging.MustGetLogger("example")

//var format = "%{color}%{time:2006-01-02T15:04:05.9999Z-07:00} → %{level:.4s} %{id:03x}%{color:reset} %{message}"
var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.9999Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")
var syslog_log_format = logging.MustStringFormatter("[%{level:.1s}] {shortpkg}[%{longfunc}] %{message}")

func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	syslogBackend, err := logging.NewSyslogBackend("zeroctl")
	if err != nil {
		log.Fatal(err)
	}
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	syslogFormatter := logging.NewBackendFormatter(syslogBackend, syslog_log_format)
	logging.SetBackend(stderrFormatter, syslogFormatter)
	logging.SetFormatter(syslog_log_format)
	var cfg map[string]interface{}
	raw_cfg, err := ioutil.ReadFile("cfg/zeroctl.conf")
	err = yaml.Unmarshal([]byte(raw_cfg), &cfg)
	d, err := yaml.Marshal(&cfg)
	log.Info("Config:")
	for _, line := range strings.Split(string(d), "\n") {
		log.Info(line)
	}
	go broadcast("224.1.2.3:54321")
	a := protocol.NewContainer()
	a.Body = []byte("asd")
	//	b := []int{1, 2, 3, 4}
	//	a.body = byte("test")
	fmt.Printf("%v\n", a)
	txt, _ := yaml.Marshal(&a)
	log.Info(string(txt))
	time.Sleep(10000 * time.Millisecond)
}

func broadcast(addr string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Info("network error")
	}
	uuid := make([]byte, 32)
	hostname, _ := os.Hostname()
	rand.Read(uuid)
	node := protocol.NewNode(hostname, uuid)
	transport := amqp09.NewTransport()
	for {
		packet := protocol.NewContainer()
		hb := node.NewHeartbeat()
		packet.Body, _ = yaml.Marshal(hb.Headers)
		transport.SendEvent(hb, "discovery.service", "cake")
		log.Debug("Sent hb")

		fmt.Fprintf(conn, string(packet.Body))
		time.Sleep(1000 * time.Millisecond)
	}
}

//func
