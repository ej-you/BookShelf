package main

import (
	"log"

	"BookShelf/config"
)

func main() {
	// load config
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config: %+v", cfg)

	// TODO: run app
}
