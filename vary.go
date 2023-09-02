package vary

import "reflect"

type Interfaces map[ID]*Interface

// NewInterfaces registry.
func NewInterfaces() Interfaces {
	return Interfaces{}
}

// New interface.
func (r Interfaces) New(v interface{}) *Interface {
	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Interface {
		panic("must be an pointer to interface")
	}

	i := &Interface{Self: t.Elem(), Implementations: []reflect.Type{}}

	r[i.ID()] = i

	return i
}

type Interface struct {
	Self            reflect.Type
	Implementations []reflect.Type
}

func (i *Interface) ID() ID {
	return NewID(i.Self)
}

func (i *Interface) Add(v interface{}) struct{} {
	t := reflect.TypeOf(v)

	if !t.Implements(i.Self) {
		panic("type does not implement interface: " + i.ID())
	}

	i.Implementations = append(i.Implementations, t)

	return struct{}{}
}

// ID is a unique identifier for a type.
type ID string

// NewID for the type.
func NewID(t reflect.Type) ID {
	if t == nil {
		return ""
	}
	return ID(t.PkgPath() + "." + t.Name())
}
