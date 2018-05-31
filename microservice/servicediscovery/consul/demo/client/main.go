package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	services, qm, err := client.Catalog().Service("web", "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("qm: %v", qm)
	for _, v := range services {
		fmt.Printf("services: %v", v)
	}
	fmt.Println("")
	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "foo", Value: []byte("test")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v", pair)
}
