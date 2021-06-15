package main

import (
	"log"

	"fmt"
)

type ModuleInit func() (*Module, error)

var moduleRegistry = map[string]ModuleInit{}

func RegisterModule(typ string, initHandler ModuleInit) {
	_, exists := moduleRegistry[typ]
	if exists {
		log.Fatalf("module %v has already been registered", typ)
	}
	moduleRegistry[typ] = initHandler
}

func GetModule(typ string) (*Module, error) {
	init, exists := moduleRegistry[typ]
	if exists {
		return init()
	}
	return nil, fmt.Errorf("module %v not found, check your configuration", typ)
}
