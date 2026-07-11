package gnmi

import (
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/pkg/api/path"
)

func BuildPath(p string) (*gnmi.Path, error) {
	return path.ParsePath(p)
}

func MustBuildPath(p string) *gnmi.Path {
	path, err := path.ParsePath(p)
	if err != nil {
		panic("invalid gnmi path: " + p + ": " + err.Error())
	}
	return path
}
