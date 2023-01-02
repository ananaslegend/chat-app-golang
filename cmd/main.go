package main

import (
	"chat-app-golang/internal/app"
	"log"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		log.Println(err)
	}
}
