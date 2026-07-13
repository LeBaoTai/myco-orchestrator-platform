package main

import (
	"fmt"
	"log"

	"github.com/LeBaoTai/myco-controller/internal/config"
	"github.com/LeBaoTai/myco-controller/internal/schema"
)

func main() {
	cfg, err := config.Load("./internal/config/schema.yml")
	if err != nil {
		log.Fatal(err)
	}

	loader := schema.NewYangLoader(
		schema.LoaderConfig{
			Repository:  cfg.Schema.Repository,
			SearchPaths: cfg.Schema.SearchPaths,
			Modules:     cfg.Schema.Modules,
		},
	)

	fmt.Printf("Loader:%v\n", loader)

	reg, err := loader.Load()

	fmt.Printf("Reg: %v\n", reg)
}
