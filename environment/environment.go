package environment

type Environment struct {
	store map[string]any
	paren *Environment
}

func New(parent *Environment) *Environment {
	s := make(map[string]any)
	if parent == nil {
		parent = &Environment{
			store: make(map[string]any),
		}
		parent.store["nil"] = nil
		parent.store["true"] = true
		parent.store["false"] = false
		parent.store["pi"] = 3.14159265358979323846264338327950288419716939937510582097494459
		parent.store["e"] = 2.71828182845904523536028747135266249775724709369995957496696763
	}
	return &Environment{
		store: s,
		paren: parent,
	}
}

func (e *Environment) Set(name string, val any) {
	e.store[name] = val
}

func (e *Environment) Get(name string) (any, bool) {
	val, ok := e.store[name]
	if ok {
		return val, ok
	}
	if e.paren != nil {
		return e.paren.Get(name)
	}
	return nil, false
}

func (e *Environment) Delete(name string) {
	delete(e.store, name)
}
