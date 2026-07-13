package utils

import "google.golang.org/grpc/metadata"

func CreateMetadata(username, password string) metadata.MD {
	md := metadata.Pairs("username", username, "password", password)
	return md
}
