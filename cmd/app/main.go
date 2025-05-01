// App binary starts web application.
package main

import (
	"log"

	"BookShelf/config"
	"BookShelf/internal/app"
)

func main() {
	if err := startApp(); err != nil {
		log.Fatal(err)
	}
}

func startApp() error {
	// load config
	cfg, err := config.New()
	if err != nil {
		return err
	}
	// create app
	application, err := app.New(cfg)
	if err != nil {
		return err
	}
	// run app
	if err := application.Run(); err != nil {
		return err
	}
	return nil
}
