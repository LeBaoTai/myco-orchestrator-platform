package schema

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/openconfig/goyang/pkg/yang"
)

type YangLoader struct {
	ldCfg *LoaderConfig
}

func NewYangLoader(
	cfg LoaderConfig,
) *YangLoader {
	return &YangLoader{
		ldCfg: &cfg,
	}
}

func (l *YangLoader) Load() (*Registry, error) {

	log.Println("Loader..")

	modules := yang.NewModules()

	// TODO: add search paths
	for _, p := range l.ldCfg.SearchPaths {
		full := filepath.Join(l.ldCfg.Repository, p)

		modules.Path = append(modules.Path, full)
	}

	// TODO: read modules
	for _, m := range l.ldCfg.Modules {
		file := m + ".yang"
		if err := modules.Read(file); err != nil {
			return nil, err
		}
	}

	// TODO: process
	if errs := modules.Process(); len(errs) > 0 {
		return nil, fmt.Errorf("%v", errs)
	}

	// TODO: build registry
	registry := NewRegistry()

	for name, module := range modules.Modules {
		entry := yang.ToEntry(module)

		registry.RegisterModule(name, entry)
	}

	log.Println(l.ldCfg.Repository)

	log.Println(l.ldCfg.SearchPaths)

	log.Println(l.ldCfg.Modules)

	// TODO: walk(entry)

	return registry, nil
}
