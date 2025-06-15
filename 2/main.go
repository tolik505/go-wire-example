package main

import (
	"log"
)

func main() {
	app, err := initApp(ConnectionInfo("info"))
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
