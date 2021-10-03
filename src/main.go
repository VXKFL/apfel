package main

import (
	"apfel/database"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
