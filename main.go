package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	addr := "192.100.100.101:57400"
	username := "admin"
	password := "NokiaSrl1!"

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	md := metadata.Pairs("username", username, "password", password)
	ctx = metadata.NewOutgoingContext(ctx, md)

	client := gnmi.NewGNMIClient(conn)

	req := &gnmi.GetRequest{
		Path: []*gnmi.Path{
			{
				Elem: []*gnmi.PathElem{
					{Name: "system"},
				},
			},
		},

		Type:     gnmi.GetRequest_STATE,
		Encoding: gnmi.Encoding_JSON_IETF,
	}

	res, err := client.Get(ctx, req)
	if err != nil {
		log.Printf("Failed to get: %v", err)
		return
	}

	fmt.Printf("Admin's password: %v\n", res.String())
}
