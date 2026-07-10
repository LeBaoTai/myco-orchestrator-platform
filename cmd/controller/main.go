package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeBaoTai/myco-controller/internal/controller/gnmi"
)

func main() {

	client, err := gnmi.New(gnmi.Config{
		Address:  "192.100.100.101:57400",
		Username: "admin",
		Password: "NokiaSrl1!",
		Timeout:  5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hostnamePath, err := gnmi.BuildPath("/system/state/hostname")
	if err != nil {
		log.Printf("failed to build path: %v", err)
	}

	res, err := client.Get(hostnamePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
