package vary_test

import (
	"reflect"
	"testing"

	"github.com/ysmood/got"
	"github.com/ysmood/vary"
)

type A interface {
	T()
}

var iA = vary.New(new(A), &C{})

type B struct{}

func (b B) T() {}

var _ = iA.Add(B{})

type C struct{}

func (c *C) T() {}

func TestNew(t *testing.T) {
	g := got.T(t)

	g.Len(vary.Get(iA.ID()).Implementations, 2)
}

func TestID(t *testing.T) {
	g := got.T(t)

	g.Eq(vary.ID(reflect.TypeOf(nil)), "")
	g.Eq(vary.ID(reflect.TypeOf(1)), ".int")
	g.True(iA.Has(B{}))
	g.True(iA.Has(&B{}))
	g.True(iA.Has(&C{}))
	g.False(iA.Has(C{}))
}

func TestCollision(t *testing.T) {
	g := got.T(t)

	g.Eq(g.Panic(func() {
		vary.New(new(A), &C{})
	}), "interface already registered: github.com/ysmood/vary_test.A")
}
