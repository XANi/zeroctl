package protocol
import (
	//	"encoding/base64"
	"fmt"
)
type Event struct {
	Headers map[string]string
	Body []byte
}

func (node Node) NewEvent() (r Event) {
	r.Headers = make(map[string]string)
	r.Headers[`node-uuid`] = fmt.Sprintf(`%x`,node.UUID)
	r.Headers[`node-name`] = string(node.Name)
	return r
}
