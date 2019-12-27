// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node_test

import (
	"testing"

	"github.com/go-trellis/node"

	. "github.com/smartystreets/goconvey/convey"
)

const runTimes int = 10000

var mapRunTimes map[string]int

func init() {
	mapRunTimes = make(map[string]int)
}

func TestRandom(t *testing.T) {

	Convey("get nil", t, func() {
		rNil := node.New(node.NodeTypeRandom, "")
		r := node.New(node.NodeTypeRandom, "random")
		Convey("when not initial nodes", func() {
			r.Add(nil)
			Convey("will return nil", func() {
				So(rNil, ShouldBeNil)
				So(r.IsEmpty(), ShouldBeTrue)
				node, ok := r.NodeFor()
				So(node, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})

		Convey("when initial normal nodes and remove nodes", func() {
			mapM, err := node.NewNodesFromConfig("sample.json")
			So(err, ShouldBeNil)
			r = mapM["random_test1"]
			r.Remove()
			Convey("will return nil", func() {
				So(r.IsEmpty(), ShouldBeTrue)
				node, ok := r.NodeFor()
				So(node, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
		Convey("when initial normal nodes and remove all nodes by IDs", func() {
			mapM, err := node.NewNodesFromConfig("sample.json")
			So(err, ShouldBeNil)
			r = mapM["random_test1"]
			Convey("will return normal node", func() {
				So(r.IsEmpty(), ShouldBeFalse)
				node, ok := r.NodeFor()
				So(node, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			})
			r.RemoveByID("node2")
			r.RemoveByID("node2")
			r.RemoveByID("node1")
			r.RemoveByID("node1")
			Convey("will return nil", func() {
				So(r.IsEmpty(), ShouldBeTrue)
				node, ok := r.NodeFor()
				So(node, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
		Convey("add new nodes", func() {
			r.Add(&node.Node{
				ID:     "1",
				Weight: 2,
				Value:  "test1",
			})
			Convey("will return normal node", func() {
				So(r.IsEmpty(), ShouldBeFalse)
				node, ok := r.NodeFor()
				So(node, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("get normal node", t, func() {
		mapM, err := node.NewNodesFromConfig("sample.json")
		So(err, ShouldBeNil)
		r := mapM["random_test2"]
		Convey("test run times", func() {
			Convey("node1:node2 vnode_number 20:80", func() {
				for i := 0; i < runTimes; i++ {
					node, _ := r.NodeFor("")
					mapRunTimes[node.ID]++
				}

				t.Log(mapRunTimes["1"])
				t.Log(mapRunTimes["2"])

				So(mapRunTimes["1"], ShouldBeBetweenOrEqual, 1850, 2150)
				So(mapRunTimes["2"], ShouldBeBetweenOrEqual, 7850, 8150)
			})
		})

	})
}
