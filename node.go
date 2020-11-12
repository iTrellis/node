// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"

	"github.com/go-trellis/config"
)

// Type define node type
type Type uint8

// NodeType
const (
	NodeTypeDirect Type = iota
	NodeTypeRandom
	NodeTypeConsistent
	NodeTypeRoundRobin
)

// Node params for a node
type Node struct {
	// for recognize node with input id
	ID string `yaml:"id" json:"id"`
	// node's probability weight, roundrobin does not support
	Weight uint32 `yaml:"weight" json:"weight"`
	// node's value
	Value string `yaml:"value" json:"value"`
	// kvs for meta data
	Metadata config.Options `yaml:"options" json:"options"`

	number uint32
}

// Manager node manager functions defines.
type Manager interface {
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

// New get node manager by node type
func New(nt Type, name string) (Manager, error) {
	switch nt {
	case NodeTypeDirect:
		return NewDirect(name)
	case NodeTypeRandom:
		return NewRadmon(name)
	case NodeTypeConsistent:
		return NewConsistent(name)
	case NodeTypeRoundRobin:
		return NewRoundRobin(name)
	default:
		return nil, fmt.Errorf("not supperted type: %d", nt)
	}
}
