package vary

import "reflect"

type Interfaces map[TypeID]*Interface

// NewInterfaces registry.
func NewInterfaces() Interfaces {
	return Interfaces{}
}

// New interface.
func (r Interfaces) New(v interface{}, vs ...interface{}) *Interface {
	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Interface {
		panic("must be an pointer to interface")
	}

	i := &Interface{Self: t.Elem(), Implementations: map[TypeID]reflect.Type{}}

	if _, has := r[i.ID()]; has {
		panic("interface already registered: " + i.ID())
	}

	r[i.ID()] = i

	for _, v := range vs {
		i.Add(v)
	}

	return i
}

type Interface struct {
	Self            reflect.Type
	Implementations map[TypeID]reflect.Type
}

func (i *Interface) ID() TypeID {
	return ID(i.Self)
}

func (i *Interface) Add(v interface{}) TypeID {
	t := reflect.TypeOf(v)

	if !t.Implements(i.Self) {
		panic("type does not implement interface: " + i.ID())
	}

	id := ID(t)
	i.Implementations[id] = t

	return id
}

// Has returns true if the v has bind to i.
func (i *Interface) Has(v interface{}) bool {
	t := reflect.TypeOf(v)
	_, has := i.Implementations[ID(t)]
	return has
}

// TypeID is a unique identifier for a type.
type TypeID string

// ID for the type.
func ID(t reflect.Type) TypeID {
	if t == nil {
		return ""
	}
	return TypeID(t.PkgPath() + "." + t.Name())
}
