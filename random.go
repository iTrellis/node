// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type radmon struct {
	Name string

	nodes map[string]*Node
	rings map[int64]*Node

	count int64

	sync.RWMutex
}

// NewRadmon get random node manager
func NewRadmon(name string) (Manager, error) {
	if name = strings.TrimSpace(name); name == "" {
		return nil, fmt.Errorf("name should not be nil")
	}
	return &radmon{Name: name}, nil
}

func (p *radmon) IsEmpty() bool {
	return atomic.LoadInt64(&p.count) == 0
}

func (p *radmon) Add(node *Node) {
	if node == nil {
		return
	}
	p.Lock()
	defer p.Unlock()
	p.add(node)
}

func (p *radmon) add(pNode *Node) {
	if p.nodes == nil {
		p.nodes = make(map[string]*Node)
	}

	p.nodes[pNode.ID] = pNode

	p.updateRings()
}

func (p *radmon) Remove() {
	p.Lock()
	defer p.Unlock()
	p.remove()
}

func (p *radmon) remove() {
	p.nodes = nil
	p.rings = nil
	p.count = 0
}

func (p *radmon) RemoveByID(id string) {
	p.Lock()
	defer p.Unlock()
	p.removeByID(id)
}

func (p *radmon) removeByID(id string) {
	if p.nodes == nil {
		return
	} else if p.IsEmpty() {
		p.remove()
		return
	}

	node := p.nodes[id]
	if node == nil {
		return
	}

	delete(p.nodes, id)
	p.updateRings()
}

func (p *radmon) NodeFor(...string) (*Node, bool) {
	p.RLock()
	defer p.RUnlock()
	if p.IsEmpty() {
		return nil, false
	}

	rand.Seed(time.Now().UnixNano())

	return p.rings[rand.Int63n(p.count)], true
}

func (p *radmon) updateRings() {
	p.rings = make(map[int64]*Node)

	p.count = 0
	for _, v := range p.nodes {

		for i := 0; i < int(v.Weight); i++ {
			ring := *v
			ring.number = uint32(i + 1)
			p.rings[p.count] = &ring

			p.count++
		}
	}
}

func (p *radmon) PrintNodes() {
	p.RLock()
	defer p.RUnlock()

	for i, v := range p.nodes {
		fmt.Println("nodes:", i, *v)
	}
}
