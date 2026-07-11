package gnmi

import "github.com/openconfig/gnmi/proto/gnmi"

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
