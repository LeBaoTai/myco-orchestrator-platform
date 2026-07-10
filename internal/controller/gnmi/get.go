package gnmi

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

// Get lấy giá trị tại các path chỉ định, mặc định lấy STATE
func (c *Client) Get(paths ...*gnmi.Path) (*gnmi.GetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()

	req := &gnmi.GetRequest{
		Path:     paths,
		Type:     gnmi.GetRequest_STATE,
		Encoding: gnmi.Encoding_JSON_IETF,
	}
	return c.gnmiC.Get(ctx, req)
}

// GetConfig lấy giá trị config (thay vì state)
func (c *Client) GetConfig(paths ...*gnmi.Path) (*gnmi.GetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()

	req := &gnmi.GetRequest{
		Path:     paths,
		Type:     gnmi.GetRequest_CONFIG,
		Encoding: gnmi.Encoding_JSON_IETF,
	}
	return c.gnmiC.Get(ctx, req)
}

func (c *Client) GetString(path string) (string, error) {
	p, err := BuildPath(path)
	if err != nil {
		return "", err
	}

	res, err := c.Get(p)
	if err != nil {
		return "", err
	}

	for _, notif := range res.GetNotification() {
		for _, upd := range notif.GetUpdate() {
			if upd.GetVal() == nil {
				continue
			}

			jsonBytes := upd.GetVal().GetJsonIetfVal()
			if len(jsonBytes) == 0 {
				continue
			}

			return string(jsonBytes), nil
		}
	}
	return "", nil
}
