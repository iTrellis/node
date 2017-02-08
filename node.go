// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package node

type NodeType uint8

const (
	NodeTypeDirect NodeType = iota
	NodeTypeRandom
	NodeTypeConsistent
)

type Node struct {
	Id     string
	Weight uint32
	Value  interface{}
	number uint32
}

type NodeManager interface {
	// adds a node to the node ring.
	Add(node *Node)
	// get the node responsible for the data key.
	NodeFor(keys ...string) (*Node, bool)
	// removes all nodes from the node ring.
	Remove()
	// removes a node from the node ring.
	RemoveByID(id string)
	// print all nodes
	PrintNodes()
	// is the node ring empty
	IsEmpty() bool
}

func New(nt NodeType, name string) NodeManager {
	switch nt {
	case NodeTypeDirect:
		return NewDirect(name)
	case NodeTypeRandom:
		return NewRadmon(name)
	case NodeTypeConsistent:
		return NewConsistent(name)
	default:
		return nil
	}
}
