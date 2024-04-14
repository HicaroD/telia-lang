package scope

import (
	"errors"
	"fmt"
)

var (
	SYMBOL_ALREADY_DEFINED_ON_SCOPE error = errors.New("symbol already defined on scope")
	SYMBOL_NOT_FOUND_ON_SCOPE       error = errors.New("symbol not found on scope")
)

type Scope[V any] struct {
	parent *Scope[V]
	nodes  map[string]V
}

func New[V any](parent *Scope[V]) *Scope[V] {
	return &Scope[V]{parent: parent, nodes: map[string]V{}}
}

func (scope *Scope[V]) Insert(name string, element V) error {
	if _, ok := scope.nodes[name]; ok {
		return fmt.Errorf("%s: %s", SYMBOL_ALREADY_DEFINED_ON_SCOPE, name)
	}
	scope.nodes[name] = element
	return nil
}

func (scope *Scope[V]) Lookup(name string) (V, error) {
	if node, ok := scope.nodes[name]; ok {
		return node, nil
	}
	if scope.parent == nil {
		// HACK
		var empty V
		return empty, fmt.Errorf("%s: %s", SYMBOL_NOT_FOUND_ON_SCOPE, name)
	}
	return scope.parent.Lookup(name)
}

func (scope Scope[V]) String() string {
	return fmt.Sprintf("Scope:\nParent: %v\nCurrent: %v\n", scope.parent, scope.nodes)
}
