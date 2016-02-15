package main

import (
	"log"

	"github.com/Pholey/Exgo/resources"

	_ "github.com/Pholey/Exgo/db"
	_ "github.com/Pholey/Exgo/redis"
	// config "github.com/Pholey/Exgo/config"
)

func main() {
	// Start the server
	println("Listening at :8080")
	log.Fatal(resources.NewRouter().Run(":8080"))
}
