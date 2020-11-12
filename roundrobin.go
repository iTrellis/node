// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/go-trellis/common/formats"
)

type roundrobin struct {
	Name string

	nodes   map[string]*Node
	indexes formats.Strings

	count      int64
	robinIndex int64

	sync.RWMutex
}

// NewRoundRobin get roundrobin node manager
func NewRoundRobin(name string) (Manager, error) {
	if name = strings.TrimSpace(name); name == "" {
		return nil, fmt.Errorf("name should not be nil")
	}
	return &roundrobin{Name: name}, nil
}

func (p *roundrobin) IsEmpty() bool {
	return atomic.LoadInt64(&p.count) == 0
}

func (p *roundrobin) Add(node *Node) {
	if node == nil {
		return
	}

	if node.Weight != 0 {
		node.Weight = 0
	}

	p.Lock()
	defer p.Unlock()
	p.add(node)
}

func (p *roundrobin) add(pNode *Node) {
	if p.nodes == nil {
		p.nodes = make(map[string]*Node)
	}

	_, ok := p.nodes[pNode.ID]
	if !ok {
		p.count++
		p.indexes = append(p.indexes, pNode.ID)

		sort.Sort(p.indexes)
	}

	p.nodes[pNode.ID] = pNode
}

func (p *roundrobin) NodeFor(...string) (*Node, bool) {
	if p.IsEmpty() {
		return nil, false
	}
	p.RLock()
	defer p.RUnlock()

	if p.robinIndex >= p.count {
		p.robinIndex = 0
	}
	node := p.nodes[p.indexes[int(p.robinIndex%p.count)]]

	p.robinIndex++

	return node, true
}

func (p *roundrobin) Remove() {
	p.Lock()
	defer p.Unlock()
	p.remove()
}

func (p *roundrobin) remove() {
	p.nodes = nil
	p.indexes = nil
	p.count = 0
	p.robinIndex = 0
}

func (p *roundrobin) RemoveByID(id string) {
	p.Lock()
	defer p.Unlock()
	p.removeByID(id)
}

func (p *roundrobin) removeByID(id string) {
	if p.IsEmpty() {
		p.remove()
		return
	}

	_, ok := p.nodes[id]
	if !ok {
		return
	}
	delete(p.nodes, id)
	p.removeKey(id)
	p.count--

	sort.Sort(p.indexes)
}

func (p *roundrobin) removeKey(key string) {
	for i, v := range p.indexes {
		if v == key {
			p.indexes = append(p.indexes[:i], p.indexes[i+1:]...)
			break
		}
	}
}

func (p *roundrobin) PrintNodes() {
	p.RLock()
	defer p.RUnlock()

	for i, v := range p.nodes {
		fmt.Println("nodes:", i, *v)
	}
}
