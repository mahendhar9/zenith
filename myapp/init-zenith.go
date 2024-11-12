package main

import (
	"log"
	"os"

	"github.com/mahendhar9/zenith"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	zen := &zenith.Zenith{}
	err = zen.New(path)
	if err != nil {
		log.Fatal(err)
	}

	zen.AppName = "myApp"
	zen.Debug = true

	app := &application{
		App: zen,
	}

	return app

}
