package main

import (
	//	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
)

var log = logging.MustGetLogger("example")

//var format = "%{color}%{time:2006-01-02T15:04:05.9999Z-07:00} → %{level:.4s} %{id:03x}%{color:reset} %{message}"
var stdout_log_format = "%{color:bold}%{time:2006-01-02T15:04:05.9999Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}"
var syslog_log_format = "[%{level:.1s}] {shortpkg}[%{longfunc}] %{message}"

func main() {
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	//syslogBackend, err := logging.NewSyslogBackend("zeroctl")
	//if err != nil {
	//	log.Fatal(err)
	//}
	logging.SetBackend(logBackend) //, syslogBackend)
	logging.SetFormatter(logging.MustStringFormatter(stdout_log_format))
	var cfg map[string]interface{}
	raw_cfg, err := ioutil.ReadFile("cfg/zeroctl.conf")
	err = yaml.Unmarshal([]byte(raw_cfg), &cfg)
	d, err := yaml.Marshal(&cfg)
	log.Info("Config:")
	log.Info(string(d))
	_ = err // please golang just fuck off from that variable for testing

}
