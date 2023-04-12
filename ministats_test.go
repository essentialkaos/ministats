package ministats

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type MinistatsSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&MinistatsSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *MinistatsSuite) TestBasic(c *C) {
	d := NewData(3)

	c.Assert(d.Min(), Equals, uint64(0))
	c.Assert(d.Max(), Equals, uint64(0))
	c.Assert(d.Cap(), Equals, 3)
	c.Assert(d.Len(), Equals, 0)

	d.Add(1)
	d.Add(5)
	d.Add(3)

	c.Assert(d.Min(), Equals, uint64(1))
	c.Assert(d.Max(), Equals, uint64(5))
	c.Assert(d.Len(), Equals, 3)

	d.Add(2)
	d.Add(4)

	c.Assert(d.Min(), Equals, uint64(2))
	c.Assert(d.Max(), Equals, uint64(4))

	d.Reset()

	c.Assert(d.Min(), Equals, uint64(0))
	c.Assert(d.Max(), Equals, uint64(0))
}

func (s *MinistatsSuite) TestMin(c *C) {
	d := NewData(5)

	c.Assert(d.Min(), Equals, uint64(0))

	d.Add(1)
	d.Add(6)
	d.Add(12)
	d.Add(3)
	d.Add(5)

	c.Assert(d.Min(), Equals, uint64(1))
}

func (s *MinistatsSuite) TestMax(c *C) {
	d := NewData(5)

	c.Assert(d.Max(), Equals, uint64(0))

	d.Add(1)
	d.Add(6)
	d.Add(12)
	d.Add(3)
	d.Add(5)

	c.Assert(d.Max(), Equals, uint64(12))
}

func (s *MinistatsSuite) TestMean(c *C) {
	d := NewData(5)

	c.Assert(d.Mean(), Equals, uint64(0))

	d.Add(1)
	d.Add(6)
	d.Add(12)
	d.Add(3)
	d.Add(5)

	c.Assert(d.Mean(), Equals, uint64(5))
}

func (s *MinistatsSuite) TestStdDev(c *C) {
	d := NewData(5)

	c.Assert(d.StdDevS(), Equals, uint64(0))
	c.Assert(d.StdDevP(), Equals, uint64(0))

	d.Add(5)
	d.Add(10)
	d.Add(50)
	d.Add(20)
	d.Add(5)

	c.Assert(d.StdDevS(), Equals, uint64(18))
	c.Assert(d.StdDevP(), Equals, uint64(16))
}

func (s *MinistatsSuite) TestPercentile(c *C) {
	d := NewData(300)

	p := d.Percentile(1, 5, 90, 95, 99)

	c.Assert(p, HasLen, 5)
	c.Assert(p, DeepEquals, []uint64{0, 0, 0, 0, 0})

	for i := 0; i <= 280; i++ {
		d.Add(uint64(i))
	}

	p = d.Percentile(0, 1, 5, 95, 99, 99.999)

	c.Assert(p, HasLen, 6)
	c.Assert(p, DeepEquals, []uint64{0, 2, 14, 266, 278, 280})

	d.Add(300)

	p = d.Percentile(0, 1, 5, 95, 99, 99.999)

	c.Assert(p, HasLen, 6)
	c.Assert(p, DeepEquals, []uint64{0, 2, 14, 267, 279, 300})
}
