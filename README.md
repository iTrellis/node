# node

a nodes ring for a key to get node

## Build

* [![Build Status](https://travis-ci.org/go-trellis/node.png)](https://travis-ci.org/go-trellis/node)


## Usage

```go
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

## Get Node Methods

### Consistent hashing

* [WIKI](https://en.wikipedia.org/wiki/Consistent_hashing)
* [consistent:test](consistent_test.go)

### Random

* [random:test](random_test.go)