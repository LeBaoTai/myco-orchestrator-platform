package gnmi

import (
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/pkg/api/path"
)

// BuildPath convert string path (vd "system/name/host-name") sang *gnmi.Path
func BuildPath(p string) (*gnmi.Path, error) {
	return path.ParsePath(p)
}

// MustBuildPath panic nếu path sai — dùng khi path là hardcode, chắc chắn đúng
func MustBuildPath(p string) *gnmi.Path {
	path, err := path.ParsePath(p)
	if err != nil {
		panic("invalid gnmi path: " + p + ": " + err.Error())
	}
	return path
}
