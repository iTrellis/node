// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"
	"strings"
)

type direct struct {
	Name string
	node *Node
}

// NewDirect get direct node manager
func NewDirect(name string) (Manager, error) {
	if name = strings.TrimSpace(name); name == "" {
		return nil, fmt.Errorf("name should not be nil")
	}
	return &direct{Name: name}, nil
}

func (p *direct) IsEmpty() bool {
	return p.node == nil
}

func (p *direct) Add(node *Node) {
	if node == nil {
		return
	}
	p.add(node)
}

func (p *direct) add(pNode *Node) {
	p.node = pNode
}

func (p *direct) NodeFor(keys ...string) (*Node, bool) {
	if p.node == nil {
		return nil, false
	}
	node := *p.node

	return &node, true
}

func (p *direct) Remove() {
	p.remove()
}

func (p *direct) remove() {
	p.node = nil
}

func (p *direct) RemoveByID(id string) {
	p.removeByID(id)
}

func (p *direct) removeByID(id string) {
	if p.node == nil {
		return
	}
	if p.node.ID == id {
		p.node = nil
	}
}

func (p *direct) PrintNodes() {
	fmt.Println("node:", p.node)
}
