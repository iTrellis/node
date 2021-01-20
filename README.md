# node

a nodes ring for a key to get node

## Build

* [![GoDoc](http://godoc.org/github.com/iTrellis/node?status.svg)](http://godoc.org/github.com/iTrellis/node)

## Get Node Methods

### Direct

> dierct to get last added node, node's wight is unavailable


### Consistent hashing

> [WIKI](https://en.wikipedia.org/wiki/Consistent_hashing)

### Random

> to get the node by random


## Usage

```go

// Node params for a node
type Node struct {
	// for recognize node with input id
	ID string
	// node's probability weight
	Weight uint32
	// node's value
	Value string
	// kvs for meta data
	Metadata config.Options

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
```

### New a node manager

```go
	directNode, _ := node.New(node.NodeTypeDirect, "direct")
	randomNode, _ := node.New(node.NodeTypeRandom, "random")
	consistentNode, _ := node.New(node.NodeTypeConsistent, "consistent")
	roundrobinNode, _ := node.New(node.NodeTypeRoundRobin, "roundrobin")
```

Or 

```go
	directNode, _ := node.NewDirect("direct")
	randomNode, _ := node.NewRandom("random")
	consistentNode, _ := node.NewConsistent("consistent")
	roundrobinNode, _ := node.NewRoundRobin("roundrobin")
```

