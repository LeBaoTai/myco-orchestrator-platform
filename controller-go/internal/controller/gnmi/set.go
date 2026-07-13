package gnmi

import (
	"encoding/json"

	"github.com/openconfig/gnmi/proto/gnmi"
)

func (c *Client) SetTransaction(deletes []*gnmi.Path, updates []*gnmi.Update, replaces []*gnmi.Update) (*gnmi.SetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()

	req := &gnmi.SetRequest{
		Delete:  deletes,
		Update:  updates,
		Replace: replaces,
	}
	return c.gnmiC.Set(ctx, req)
}

func NewUpdate(p *gnmi.Path, value interface{}) (*gnmi.Update, error) {
	valBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return &gnmi.Update{
		Path: p,
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_JsonIetfVal{
				JsonIetfVal: valBytes,
			},
		},
	}, nil
}

func (c *Client) DeleteConfig(paths ...*gnmi.Path) (*gnmi.SetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()
	req := &gnmi.SetRequest{
		Delete: paths,
	}
	return c.gnmiC.Set(ctx, req)
}

func (c *Client) UpdateConfig(updates ...*gnmi.Update) (*gnmi.SetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()
	req := &gnmi.SetRequest{
		Update: updates,
	}
	return c.gnmiC.Set(ctx, req)
}

func (c *Client) ReplaceConfig(replaces ...*gnmi.Update) (*gnmi.SetResponse, error) {
	ctx, cancel := c.ctxWithAuth()
	defer cancel()
	req := &gnmi.SetRequest{
		Replace: replaces,
	}
	return c.gnmiC.Set(ctx, req)
}
