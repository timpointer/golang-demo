// Package main provides ...
package main

import "testing"

func TestConfig(t *testing.T) {

	config := NewConfig("config/conf.json")
	if config == nil {
		t.Error("config init error")
	}
}
