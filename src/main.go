package main

import (
	"apfel/database"
	"log"
)

func main() {
	log.Printf("Connecting to Database...")

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
