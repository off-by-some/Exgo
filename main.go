package main

import (
  "log"
  _     "Exgo/db"
  _  "Exgo/redis"
  http   "net/http"
  resources "Exgo/resources"
  // config "Exgo/config"
)

func main() {
    // Start the server
    println("Listening at :8080")
    log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))
}
