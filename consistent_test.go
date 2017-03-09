// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package node_test

import (
	"testing"

	"github.com/go-trellis/node"

	. "github.com/smartystreets/goconvey/convey"
)

const key = "key"

func TestConsistent(t *testing.T) {

	Convey("get nil", t, func() {
		cNil := node.NewConsistent("")
		c := node.NewConsistent("test")
		Convey("when not initial nodes", func() {
			Convey("will return nil", func() {
				So(cNil, ShouldBeNil)

				So(c.IsEmpty(), ShouldBeTrue)
				value, ok := c.NodeFor(key)
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})

		Convey("when initial nodes and remove nodes", func() {
			c.Add(nil)
			c.Add(&node.Node{
				ID:     "1",
				Weight: 1,
				Value:  "test1",
			})
			c.Add(&node.Node{
				ID:     "1",
				Weight: 2,
				Value:  "test1",
			})
			c.Remove()
			Convey("will return nil", func() {
				So(c.IsEmpty(), ShouldBeTrue)
				value, ok := c.NodeFor(key)
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
		Convey("when initial nodes and remove all nodes by IDs", func() {
			c.Add(&node.Node{
				ID:     "1",
				Weight: 1,
				Value:  "test1",
			})
			c.Add(&node.Node{
				ID:     "2",
				Weight: 2,
				Value:  "test2",
			})
			Convey("will return normal node", func() {
				So(c.IsEmpty(), ShouldBeFalse)
				value, ok := c.NodeFor(key)
				So(value, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			})
			c.RemoveByID("2")
			c.RemoveByID("2")
			c.RemoveByID("1")
			c.RemoveByID("1")
			c.RemoveByID("1")
			Convey("will return nil", func() {
				So(c.IsEmpty(), ShouldBeTrue)
				value, ok := c.NodeFor(key)
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("get normal node", t, func() {
		c := node.New(node.NodeTypeConsistent, "test")
		c.Add(&node.Node{
			ID:     "1",
			Weight: 2,
			Value:  "test1",
		})
		c.Add(&node.Node{
			ID:     "2",
			Weight: 10,
			Value:  "test2",
		})

		Convey("when initial normal configure", func() {
			Convey("will return normal node", func() {
				So(c.IsEmpty(), ShouldBeFalse)
				value, ok := c.NodeFor("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				So(value, ShouldNotBeNil)
				So(ok, ShouldBeTrue)
			})

			Convey("get my test key", func() {
				c.PrintNodes()
				value, ok := c.NodeFor("my test1")
				So(value.Value, ShouldEqual, "test1")
				So(ok, ShouldBeTrue)

				value, ok = c.NodeFor("my test2")
				So(value.Value, ShouldEqual, "test2")
				So(ok, ShouldBeTrue)
			})
		})

		Convey("when remove nodes by 1", func() {
			c.RemoveByID("1")
			Convey("will return node 2", func() {
				So(c.IsEmpty(), ShouldBeFalse)

				value, ok := c.NodeFor("my test1")
				So(value.Value, ShouldEqual, "test2")
				So(ok, ShouldBeTrue)

				value, ok = c.NodeFor("my test2")
				So(value.Value, ShouldEqual, "test2")
				So(ok, ShouldBeTrue)

				value, ok = c.NodeFor("my test3")
				So(value.Value, ShouldEqual, "test2")
				So(ok, ShouldBeTrue)
			})
		})
	})
}
