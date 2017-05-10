// Package main provides ...
package main

import (
	"encoding/json"
	"log"
	"os"
)

type keyMap struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}
type provider struct {
	Github keyMap `json:"github"`
}

type site struct {
	SecurityKey string `json:"securityKey"`
	Host        string `json:"host"`
}

type configration struct {
	Auth provider `json:"auth"`
	Site site     `json:"site"`
}

func NewConfig(path string) *configration {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("the file is not exist", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := configration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &conf
}
