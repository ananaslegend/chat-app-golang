package main

import (
	"chat-app-golang/internal/app"
	"log"
)

func main() {
	var mongoConnection = "mongodb+srv://admin:adminpass@freemongocluster.ynduakg.mongodb.net/?retryWrites=true&w=majority"

	a, err := app.New(mongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
