package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LeBaoTai/myco-controller/internal/controller/gnmib"
	"github.com/LeBaoTai/myco-controller/internal/oc"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"go.yaml.in/yaml/v4"
)

type ConfigPath struct {
	Paths []string `yaml:"paths"`
}

func main() {
	pathFile, err := os.ReadFile("./internal/config/paths.yml")
	if err != nil {
		log.Panic(err)
	}
	var configPath ConfigPath
	err = yaml.Unmarshal(pathFile, &configPath)
	if err != nil {
		log.Fatalf("Cannot parse the yml file: %v", err)
	}

	client, err := gnmib.New(gnmib.Config{
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

	var pathList []*gnmi.Path
	for _, path := range configPath.Paths {
		builtPath, err := gnmib.BuildPath("/" + path)
		if err != nil {
			log.Printf("Cannot create path: %v\n", err)
		}
		pathList = append(pathList, builtPath)
	}

	log.Printf("Path List: %v", pathList)
	res, err := client.GetConfig(pathList)

	if err != nil {
		log.Printf("Cannot get config: %v\b", err)
	}

	schema, err := oc.Schema()
	rootSystem := &oc.Device{}
	schema.Root = rootSystem

	for _, n := range res.Notification {
		err := ytypes.UnmarshalNotifications(
			schema,
			[]*gnmi.Notification{n},
			&ytypes.IgnoreExtraFields{},
		)
		if err != nil {
			log.Println(err)
		}
	}

	err = ytypes.UnmarshalNotifications(
		schema,
		res.Notification,
		&ytypes.IgnoreExtraFields{},
	)
	if err != nil {
		log.Printf("Failed to unmarshal path data: %v", err)
	}

	jsonOpts := &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "    ",
	}

	jsonString, err := ygot.EmitJSON(rootSystem, jsonOpts)
	if err != nil {
		log.Printf("Failed:%v", err)
	}
	fmt.Printf("Current config: %v\n", jsonString)

}
