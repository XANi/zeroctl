package main

import (
	//	"fmt"
	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"net"
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
	go broadcast("224.1.2.3:54321")
	time.Sleep(10000 * time.Millisecond)
}

func broadcast(addr string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Info("network error")
	}
	for {
		log.Info("broadcasting")
		fmt.Fprintf(conn, "im alive")
		time.Sleep(1000 * time.Millisecond)
	}
}
