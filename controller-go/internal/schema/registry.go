package schema

import (
	"fmt"

	"github.com/openconfig/goyang/pkg/yang"
)

type Registry struct {
	Modules map[string]*yang.Entry
	Entries map[string]*yang.Entry
}

func NewRegistry() *Registry {
	return &Registry{
		Modules: make(map[string]*yang.Entry),
		Entries: make(map[string]*yang.Entry),
	}
}

func (r *Registry) RegisterModule(name string, entry *yang.Entry) {
	r.Modules[name] = entry
}

func (r *Registry) RegisterPath(path string, entry *yang.Entry) {
	r.Entries[path] = entry
}

func (r *Registry) Lookup(path string) (*yang.Entry, error) {
	entry, ok := r.Entries[path]
	if !ok {
		return nil, fmt.Errorf("schema path %q not found", path)
	}

	return entry, nil
}

func (r *Registry) Module(name string) (*yang.Entry, error) {
	entry, ok := r.Modules[name]
	if !ok {
		return nil, fmt.Errorf("module %q not found", name)
	}

	return entry, nil
}
