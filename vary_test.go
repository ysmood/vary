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

var iA = vary.New(new(A))

type B struct{}

func (b B) T() {}

var _ = iA.Add(B{})

type C struct{}

func (c *C) T() {}

var _ = iA.Add(&C{})

func TestNew(t *testing.T) {
	g := got.T(t)

	g.Len(vary.Get(iA.ID()).Implementations, 2)
}

func TestID(t *testing.T) {
	g := got.T(t)

	g.Eq(vary.NewID(reflect.TypeOf(nil)), "")
	g.Eq(vary.NewID(reflect.TypeOf(1)), ".int")
}
