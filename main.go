package main

import (
  "log"
  _     "github.com/Pholey/Exgo/db"
  _  "github.com/Pholey/Exgo/redis"
  http   "net/http"
  resources "github.com/Pholey/Exgo/resources"
  // config "github.com/Pholey/Exgo/config"
)

func main() {
    // Start the server
    println("Listening at :8080")
    log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))
}
