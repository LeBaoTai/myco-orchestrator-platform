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

func parseResponse(res *gnmi.GetResponse, out interface{}) error {
	if res == nil {
		return fmt.Errorf("response is nil")
	}

	for _, notif := range res.GetNotification() {
		for _, upd := range notif.GetUpdate() {
			if err := jsonUnmarshalVal(upd.GetVal(), out); err != nil {
				return fmt.Errorf("failed to unmarshal value: %w", err)
			}
		}
	}

	return nil
}
