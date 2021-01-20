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
