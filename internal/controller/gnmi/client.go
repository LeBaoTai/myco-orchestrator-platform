package gnmi

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	conn     *grpc.ClientConn
	gnmiC    gnmi.GNMIClient
	username string
	password string
	timeout  time.Duration
}

type Config struct {
	Address    string
	Username   string
	Password   string
	SkipVerify bool
	Timeout    time.Duration
}

func New(cfg Config) (*Client, error) {
	tlsConfig := &tls.Config{InsecureSkipVerify: cfg.SkipVerify}
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", cfg.Address, err)
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &Client{
		conn:     conn,
		gnmiC:    gnmi.NewGNMIClient(conn),
		username: cfg.Username,
		password: cfg.Password,
		timeout:  timeout,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ctxWithAuth() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	md := metadata.Pairs("username", c.username, "password", c.password)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, cancel
}
