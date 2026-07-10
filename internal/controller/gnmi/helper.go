// helper.go
package gnmi

import (
	"encoding/json"
	"fmt"

	gnmi "github.com/openconfig/gnmi/proto/gnmi"
)

func jsonUnmarshalVal(val *gnmi.TypedValue, out interface{}) error {
	if val == nil {
		return fmt.Errorf("value is nil")
	}
	jsonVal := val.GetJsonIetfVal()
	if jsonVal == nil {
		jsonVal = val.GetJsonVal()
	}
	return json.Unmarshal(jsonVal, out)
}
