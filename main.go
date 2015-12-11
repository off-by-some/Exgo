package main

import (
  "log"
  db     "Exgo/db"
  redis  "Exgo/redis"
  http   "net/http"
  resources "Exgo/resources"
  config "Exgo/config"
)

func main() {

    if (redis.ConnectionFailed == nil)  {
      println("Connected to redis")
    } else {
      panic(redis.ConnectionFailed)
    }

    println(config.File.GetString("example"))

    if (db.DatabaseConnFailed == nil)  {
      println("Connected to postgres")
    } else {
      panic(db.DatabaseConnFailed)
    }

    // Start the server
    println("Listening at :8080")
    log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))
}
