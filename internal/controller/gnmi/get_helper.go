package gnmi

import (
	"encoding/json"
	"fmt"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

func ParseTypedValue(val *gnmipb.TypedValue, out interface{}) error {
	if val == nil {
		return fmt.Errorf("value is nil")
	}

	switch v := val.Value.(type) {
	case *gnmipb.TypedValue_JsonIetfVal:
		return json.Unmarshal(v.JsonIetfVal, out)
	case *gnmipb.TypedValue_JsonVal:
		return json.Unmarshal(v.JsonVal, out)
	case *gnmipb.TypedValue_StringVal:
		if strPtr, ok := out.(*string); ok {
			*strPtr = v.StringVal
			return nil
		}
		return fmt.Errorf("value is string but out is not *string")
	case *gnmipb.TypedValue_IntVal:
		if intPtr, ok := out.(*int64); ok {
			*intPtr = v.IntVal
			return nil
		}
		return fmt.Errorf("value is int64 but out is not *int64")
	case *gnmipb.TypedValue_BoolVal:
		if boolPtr, ok := out.(*bool); ok {
			*boolPtr = v.BoolVal
			return nil
		}
		return fmt.Errorf("value is bool but out is not *bool")
	default:
		return fmt.Errorf("unsupported TypedValue type: %T", v)
	}
}

func ParseGetResponse(res *gnmipb.GetResponse, out interface{}) error {
	if res == nil {
		return fmt.Errorf("response is nil")
	}

	notifications := res.GetNotification()
	if len(notifications) == 0 {
		return fmt.Errorf("no notification in response")
	}

	updates := notifications[0].GetUpdate()
	if len(updates) == 0 {
		return fmt.Errorf("no update in response (path may exist but has no value)")
	}
	if len(updates) > 1 {
		return fmt.Errorf("expected exactly 1 update, got %d — use ParseGetResponseList instead", len(updates))
	}

	return ParseTypedValue(updates[0].GetVal(), out)
}

type PathValue struct {
	Path  string
	Value json.RawMessage
}

func ParseGetResponseList(res *gnmipb.GetResponse) ([]PathValue, error) {
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}

	var results []PathValue
	for _, notif := range res.GetNotification() {
		for _, upd := range notif.GetUpdate() {
			jsonVal := upd.GetVal().GetJsonIetfVal()
			if jsonVal == nil {
				jsonVal = upd.GetVal().GetJsonVal()
			}
			results = append(results, PathValue{
				Path:  PathToString(upd.GetPath()),
				Value: jsonVal,
			})
		}
	}
	return results, nil
}
