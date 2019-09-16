package main

import (
	"log"

	api "github.com/jekabolt/dotmarket/router"
)

func main() {
	s, err := api.InitServer()
	if err != nil {
		log.Fatalf("main:api.ParseConfig [%v]", err.Error())
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("main:s.Serve [%v]", err.Error())
	}
}
