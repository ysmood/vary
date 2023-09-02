package vary

// Reg is the default registry.
var Reg = NewRegistry()

// New interface on the [Reg].
func New(v interface{}) *Interface {
	return Reg.New(v)
}
