package vary

// Default is the default registry.
var Default = NewInterfaces()

// New interface on the [Default].
func New(v interface{}, vs ...interface{}) *Interface {
	return Default.New(v, vs...)
}

// Get interface from the [Default].
func Get(id TypeID) *Interface {
	return Default[id]
}
