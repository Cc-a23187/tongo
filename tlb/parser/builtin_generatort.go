package parser

import (
	"bytes"
	"fmt"
	"github.com/startfellows/tongo/boc"
	"go/format"
)

func GenerateVarUintTypes(max int) string {
	var b bytes.Buffer
	for i := 1; i <= max; i++ {
		fmt.Fprintf(&b,
			`
type VarUInteger%v big.Int

func (u VarUInteger%v) MarshalTLB(c *boc.Cell, tag string) error {
	i := big.Int(u)
    b := i.Bytes()
	err := c.WriteLimUint(len(b), %v)
	if err != nil {
		return err
	}
	return c.WriteBytes(b)
}

func (u *VarUInteger%v) UnmarshalTLB(c *boc.Cell, tag string) error {
	ln, err := c.ReadLimUint(%v)
	if err != nil {
		return err
	}
	val, err := c.ReadBigUint(int(ln) * 8)
	if err != nil {
		return err
	}
	*u = VarUInteger%v(*val)
	return nil
}
`,
			i, i, i-1, i, i-1, i)
	}
	bytes, _ := format.Source(b.Bytes())
	return string(bytes)
}

func GenerateConstantInts(max int) string {
	var b bytes.Buffer
	for i := 1; i <= max; i++ {
		p := nearestPow(i)
		if p == i {
			continue
		}
		fmt.Fprintf(&b, `
type Uint%v uint%v

func (u Uint%v) MarshalTLB(c *boc.Cell, tag string) error {
	return c.WriteUint(uint64(u), %v)
}

func (u *Uint%v) UnmarshalTLB(c *boc.Cell, tag string) error {
	v, err := c.ReadUint(%v)
	*u = Uint%v(v)
	return err
}

type Int%v int%v

func (u Int%v) MarshalTLB(c *boc.Cell, tag string) error {
	return c.WriteInt(int64(u), %v)
}

func (u *Int%v) UnmarshalTLB(c *boc.Cell, tag string) error {
	v, err := c.ReadInt(%v)
	*u = Int%v(v)
	return err
}
`, i, p, i, i, i, i, i, i, p, i, i, i, i, i)
	}
	bytes, _ := format.Source(b.Bytes())
	return string(bytes)
}

func nearestPow(i int) int {
	switch {
	case i <= 8:
		return 8
	case i <= 16:
		return 16
	case i <= 32:
		return 32
	case i <= 64:
		return 64
	default:
		panic(i)
	}
}

type Int2 int8

func (u Int2) MarshalTLB(c *boc.Cell, tag string) error {
	return c.WriteInt(int64(u), 2)
}
func (u *Int2) UnmarshalTLB(c *boc.Cell, tag string) error {
	v, err := c.ReadInt(2)
	*u = Int2(v)
	return err
}
