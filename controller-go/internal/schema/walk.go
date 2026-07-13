package schema

import (
	"log"

	"github.com/openconfig/goyang/pkg/yang"
)

func Walk(
	registry *Registry,
	entry *yang.Entry,
) {
	log.Println("Walker")
}
