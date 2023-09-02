package vary

// Default is the default registry.
var Default = NewInterfaces()

// New interface on the [Default].
func New(v interface{}) *Interface {
	return Default.New(v)
}

// Get interface from the [Default].
func Get(id ID) *Interface {
	return Default[id]
}
