package status

import (
	"time"
)

type Nodes struct {
	Nodes map[[32]byte]Node
}


// Distance measures number of hops needed to get to the target node. 0 means it is visible locally
// Path map shows paths over which node is available
// PathDistance map shows hops needed for each path
type Node struct {
	Name string
	UUID [32]byte
	Distance uint32
	Path map[[32]byte]Path
	PathDistance map[[32]byte]uint32
}

// PathUUID is generated UUID of path
// NodeUUID is UUID of node providing a path
// Prio is path priority
// LastUse was last time path was seen working
type Path struct {
	PathUUID [32]byte
	NodeUUID [32]byte
	Prio uint32
	LastUse time.Time
}


func NewNode(name string, uuid [32]byte, distance uint32) (r Node) {
	r.Name = name
	r.UUID = uuid
	r.Distance = distance
	return r
}

func NewNodes() (r Nodes) {
	r.Nodes = map[[32]byte]Node
	return r
}

func NewPath() (r Path) {
	return r
}
