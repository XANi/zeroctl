package main

import (
	"os"
	"fmt"
	"github.com/op/go-logging"
)


var log = logging.MustGetLogger("example")
var format = "%{color}%{time:15:04:05.000000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"


func main() {
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
    syslogBackend, err := logging.NewSyslogBackend("")
	if err != nil {
        log.Fatal(err)
    }
	logging.SetBackend(logBackend, syslogBackend)
    logging.SetFormatter(logging.MustStringFormatter(format))
	fmt.Printf("Hello, world\n")
	log.Error("err")

}
