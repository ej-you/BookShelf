// App binary starts web application.
package main

import (
	"log"

	"BookShelf/config"
	"BookShelf/internal/app"
)

func main() {
	// load config
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// create app
	application, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// run app
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
