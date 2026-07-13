package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeBaoTai/myco-controller/internal/controller/gnmi"
	gnmip "github.com/openconfig/gnmi/proto/gnmi"
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

	hostnamePath, err := gnmi.BuildPath("/system/config/hostname")
	if err != nil {
		log.Printf("failed to build path: %v", err)
	}

	var listUpdate []*gnmip.Update
	upd, err := gnmi.NewUpdate(hostnamePath, "sd-hehe")
	if err != nil {
		log.Fatal(err)
	}
	listUpdate = append(listUpdate, upd)

	res, err := client.SetTransaction(nil, listUpdate, nil)
	if err != nil {
		log.Fatal(err)
	}

	setResult, err := gnmi.ParseSetResponse(res)

	fmt.Printf("Set result: %v\n", setResult)
}
