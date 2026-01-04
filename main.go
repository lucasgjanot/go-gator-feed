package main

import (
	"fmt"
	"log"

	"github.com/lucasgjanot/go-gator-feed/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("janot")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err =  config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	fmt.Printf("Read config again: %+v\n", cfg)
}
