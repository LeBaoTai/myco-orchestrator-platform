package gnmib

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/gnmi/proto/gnmi"
	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

func ParseTypedValue(val *gnmipb.TypedValue, out interface{}) error {
	if val == nil {
		return fmt.Errorf("value is nil")
	}

	// Đảm bảo biến 'out' truyền vào bắt buộc phải là một con trỏ (Pointer) để có thể gán giá trị
	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Ptr || outVal.IsNil() {
		return fmt.Errorf("output parameter 'out' must be a non-nil pointer")
	}

	switch v := val.Value.(type) {
	case *gnmi.TypedValue_JsonIetfVal:
		return json.Unmarshal(v.JsonIetfVal, out)

	case *gnmi.TypedValue_JsonVal:
		return json.Unmarshal(v.JsonVal, out)

	case *gnmi.TypedValue_StringVal:
		// Nếu 'out' là con trỏ trỏ tới chuỗi string -> gán trực tiếp
		if strPtr, ok := out.(*string); ok {
			*strPtr = v.StringVal
			return nil
		}
		// Cải tiến: Nếu 'out' là struct/map nhưng thiết bị lại gửi StringVal, thử parse JSON luôn
		return json.Unmarshal([]byte(v.StringVal), out)

	case *gnmi.TypedValue_IntVal:
		// Tối ưu bằng reflect: Chấp nhận mọi loại pointer số nguyên (int, int32, int64)
		elem := outVal.Elem()
		if elem.CanInt() {
			elem.SetInt(v.IntVal)
			return nil
		}
		return fmt.Errorf("value is int64 but target type %T cannot hold an integer", out)

	case *gnmi.TypedValue_UintVal:
		// Bổ sung thêm trường hợp UintVal (gNMI rất hay dùng cho bộ đếm Counter/Uptime)
		elem := outVal.Elem()
		if elem.CanUint() {
			elem.SetUint(v.UintVal)
			return nil
		}
		return fmt.Errorf("value is uint64 but target type %T cannot hold an unsigned integer", out)

	case *gnmi.TypedValue_BoolVal:
		if boolPtr, ok := out.(*bool); ok {
			*boolPtr = v.BoolVal
			return nil
		}
		return fmt.Errorf("value is bool but out is not *bool")

	default:
		return fmt.Errorf("unsupported TypedValue type: %T", v)
	}
}

func ParseGetResponse(res *gnmipb.GetResponse, out any) error {
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
