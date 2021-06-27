package main

import (
	"assignment/internal/app"
	"log"
)

func main() {
	log.Println("API server started...")
	app.NewServer().Start()
	defer func() {
		log.Println("API server stopped")
	}()
}
