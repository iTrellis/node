// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

type direct struct {
	Name string
	node *Node

	sync.RWMutex
}

// NewDirect get direct node manager
func NewDirect(name string) Manager {
	if name == "" {
		return nil
	}
	return &direct{Name: name}
}

func (p *direct) IsEmpty() bool {
	point := unsafe.Pointer(p.node)
	return atomic.LoadPointer(&point) == nil
}

func (p *direct) Add(node *Node) {
	if node == nil {
		return
	}
	p.Lock()
	defer p.Unlock()
	p.add(node)
}

func (p *direct) add(pNode *Node) {
	p.node = pNode
}

func (p *direct) NodeFor(keys ...string) (*Node, bool) {
	if p.IsEmpty() {
		return nil, false
	}
	p.RLock()
	defer p.RUnlock()

	return p.node, true
}

func (p *direct) Remove() {
	p.Lock()
	defer p.Unlock()
	p.remove()
}

func (p *direct) remove() {
	p.node = nil
}

func (p *direct) RemoveByID(id string) {
	p.Lock()
	defer p.Unlock()
	p.removeByID(id)
}

func (p *direct) removeByID(id string) {
	if p.IsEmpty() {
		return
	}
	if p.node.ID == id {
		p.node = nil
	}
}

func (p *direct) PrintNodes() {
	p.RLock()
	defer p.RUnlock()

	fmt.Println("node:", p.node)
}
