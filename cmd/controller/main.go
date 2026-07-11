package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeBaoTai/myco-controller/internal/controller/gnmi"
)

func main() {

	client, err := gnmi.New(gnmi.Config{
		Address:    "192.100.100.102:57400",
		Username:   "admin",
		Password:   "NokiaSrl1!",
		Timeout:    5 * time.Second,
		SkipVerify: true,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hostnamePath, err := gnmi.BuildPath("/acl/interfaces/interface[id=*]/interface-ref/state/subinterface")
	if err != nil {
		log.Printf("failed to build path: %v", err)
	}

	res, err := client.Get(hostnamePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
