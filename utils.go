package vary

// Reg is the default registry.
var interfaces = NewInterfaces()

// New interface on the default interface registry.
func New(v interface{}) *Interface {
	return interfaces.New(v)
}

// Get interface from the default interface registry.
func Get(id ID) *Interface {
	return interfaces[id]
}
