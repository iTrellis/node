/*
Copyright © 2017 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package node

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/iTrellis/common/formats"
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
