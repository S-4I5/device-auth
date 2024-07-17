package scope

type MethodMap interface {
	GetMethodScope(methodName string) (Scope, bool)
}

type methodMap struct {
	methodScopes map[string]Scope
}

func NewMapBasedMethodMap(methodScopes map[string]Scope) *methodMap {
	return &methodMap{methodScopes: methodScopes}
}

func (m *methodMap) GetMethodScope(methodName string) (Scope, bool) {
	scope, ok := m.methodScopes[methodName]
	return scope, ok
}
