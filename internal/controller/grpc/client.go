package grpc

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCClient struct {
}

type GRPClientModel struct {
	Address  string
	Username string
	Password string
}

func NewGRPCConnection() *GRPCClient {
	return &GRPCClient{}
}

func (g *GRPCClient) ConnectionToNewClient(model GRPClientModel) (*grpc.ClientConn, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(model.Address, grpc.WithTransportCredentials(creds))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return conn, nil
}
