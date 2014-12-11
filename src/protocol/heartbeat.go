package protocol

import (
	"time"
	"os"
	"fmt"
)

type Heartbeat struct {
	Headers map[string]string
}

func NewHeartbeat() (r Heartbeat) {
	r.Headers = make(map[string]string)
	r.Headers["ts"] = fmt.Sprintf("%v",time.Now().UnixNano() )
	r.Headers["hostname"], _ = os.Hostname()
	return r
}
