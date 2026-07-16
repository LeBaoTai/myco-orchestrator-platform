package gnmib

import (
	"fmt"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

type SetResult struct {
	Path    string
	Op      string
	Success bool
	Message string
}

func ParseSetResponse(res *gnmipb.SetResponse) ([]SetResult, error) {
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}

	var results []SetResult
	for _, r := range res.GetResponse() {
		result := SetResult{
			Path:    PathToString(r.GetPath()),
			Op:      r.GetOp().String(),
			Success: true,
		}

		results = append(results, result)
	}

	return results, nil
}

func CheckSetResponse(res *gnmipb.SetResponse) error {
	results, err := ParseSetResponse(res)
	if err != nil {
		return err
	}

	for _, r := range results {
		if !r.Success {
			return fmt.Errorf("set failed for path %s: %s", r.Path, r.Message)
		}
	}
	return nil
}
