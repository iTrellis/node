// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package node_test

import (
	"testing"

	"github.com/go-trellis/node"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDirect(t *testing.T) {

	Convey("get nil", t, func() {
		dNil := node.New(node.NodeTypeDirect, "")
		d := node.NewDirect("direct")
		Convey("when not initial nodes", func() {
			Convey("will return nil", func() {
				So(dNil, ShouldBeNil)

				So(d.IsEmpty(), ShouldBeTrue)
				value, ok := d.NodeFor()
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})

		Convey("when initial nodes and remove nodes", func() {
			d.Add(nil)
			d.Add(&node.Node{
				ID:     "1",
				Weight: 1,
				Value:  "test1",
			})
			d.Add(&node.Node{
				ID:     "1",
				Weight: 2,
				Value:  "test1",
			})
			d.Remove()
			Convey("will return nil", func() {
				So(d.IsEmpty(), ShouldBeTrue)
				value, ok := d.NodeFor(key)
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
		Convey("when initial nodes and remove all nodes by IDs", func() {
			d.Add(&node.Node{
				ID:     "1",
				Weight: 1,
				Value:  "test1",
			})
			d.Add(&node.Node{
				ID:     "2",
				Weight: 2,
				Value:  "test2",
			})
			Convey("will return normal node", func() {
				So(d.IsEmpty(), ShouldBeFalse)
				value, ok := d.NodeFor(key)
				So(value.ID, ShouldEqual, "2")
				So(value.Value, ShouldEqual, "test2")
				So(ok, ShouldBeTrue)
			})
			d.RemoveByID("2")
			d.RemoveByID("2")
			d.RemoveByID("1")
			d.RemoveByID("1")
			d.RemoveByID("1")
			Convey("will return nil", func() {
				So(d.IsEmpty(), ShouldBeTrue)
				value, ok := d.NodeFor(key)
				So(value, ShouldBeNil)
				So(ok, ShouldBeFalse)
			})
		})
	})
}
