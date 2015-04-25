package protocol

type Node struct {
	Name []byte
	UUID []byte
}


func NewNode(NodeName string, NodeUUID []byte) (r Node) {
	// TODO validate me!
	r.Name = []byte(NodeName)
	r.UUID = NodeUUID
	return r
}
