package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeBaoTai/myco-controller/internal/controller/gnmi"
)

func main() {
	client, err := gnmi.New(gnmi.Config{
		Address:    "192.100.100.101: 57400",
		Username:   "admin",
		Password:   "NokiaSrl1!",
		SkipVerify: true,
		Timeout:    5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	defer client.Close()

	rootPath, err := gnmi.BuildPath("/")
	if err != nil {
		log.Printf("Cannot create path: %v\n", err)
	}

	res, err := client.GetState(rootPath)
	if err != nil {
		log.Printf("Cannot get config: %v\b", err)
	}

	fmt.Printf("Current config: %v\n", res)
}
