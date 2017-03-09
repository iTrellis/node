# node

a nodes ring for a key to get node

## Build

* [![Build Status](https://travis-ci.org/go-trellis/node.png)](https://travis-ci.org/go-trellis/node)
* [![GoDoc](http://godoc.org/github.com/go-trellis/node?status.svg)](http://godoc.org/github.com/go-trellis/node)

## Get Node Methods

### Direct

dierct to get last added node, node's wight is unavailable


### Consistent hashing

* [WIKI](https://en.wikipedia.org/wiki/Consistent_hashing)
* [consistent:test](consistent_test.go)

### Random

* [random:test](random_test.go)


## Usage

```go
type Node struct {
	// for recognize node with input id
	ID string
	// node's probability weight
	Weight uint32
	// node's value
	Value interface{}
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
```

### New a node manager

```go
	directNode := node.New(node.NodeTypeDirect, "direct")
	randomNode := node.New(node.NodeTypeRandom, "random")
	consistentNode := node.New(node.NodeTypeConsistent, "consistent")
```

Or 

```go
	directNode := node.NewDirect("direct")
	randomNode := node.NewRandom("random")
	consistentNode := node.NewConsistent("consistent")
```

