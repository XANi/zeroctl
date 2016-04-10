package main

import (
	//	"fmt"
	//	"crypto/rand"
	//	"fmt"
	"github.com/op/go-logging"
	"github.com/zerosvc/go-zerosvc"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	//	"net"
	"os"
	"strings"
	"time"
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
	hostname, _ := os.Hostname()
	node := zerosvc.NewNode(hostname)
	amqpAddr := "amqp://guest:guest@localhost:5672"
	if os.Getenv(`ZEROCTL_ADDR`) != `` {
		amqpAddr = os.Getenv(`ZEROCTL_ADDR`)
	}
	var trCfg interface{}
	transport := zerosvc.NewTransport(zerosvc.TransportAMQP, amqpAddr, trCfg)
	conn_err := transport.Connect()
	if conn_err != nil {
		log.Error("Can't connect to default transport on [%s]: %+v", amqpAddr, conn_err)
		os.Exit(1)
	}
	go broadcast(node, transport)
	for {
		time.Sleep(10000 * time.Millisecond)
	}
}

func broadcast(node zerosvc.Node, transport zerosvc.Transport) {
	_ = node
	for {
		log.Debug("Sending heartbeat")
		hb := node.NewHeartbeat()
		hb.Prepare()
		transport.SendEvent(`/discovery.node`, hb)
		time.Sleep(time.Second * 3)
	}
}

//func
