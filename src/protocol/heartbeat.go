package protocol

import (
	"time"
	"os"
	"fmt"
)


func (n Node) NewHeartbeat() (r Event) {
	r  = n.NewEvent()
	r.Headers["ts"] = fmt.Sprintf("%f", float64( time.Now().UnixNano()) / 1000000000)
	//r.Headers["ts"] = fmt.Sprintf("%f",float32(time.Now().Unix()/1.000000000) )
	r.Headers["hostname"], _ = os.Hostname()
	return r
}
