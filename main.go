package main

import (
  db     "Exgo/db"
  redis  "Exgo/redis"
  http   "net/http"
  resources "Exgo/resources"
)

func main() {
    // Set up routes/handlers
    for route, handler := range resources.RoutesMap {
      resources.Router.HandleFunc(route, handler)
    }

    if (redis.ConnectionFailed == nil)  {
      println("Connected to redis")
    } else {
      panic(redis.ConnectionFailed)
    }

    if (db.DatabaseConnFailed == nil)  {
      println("Connected to postgres")
    } else {
      panic(db.DatabaseConnFailed)
    }

    // Start the server
    defer http.ListenAndServe(":8080", resources.Router)
    println("Listening at :8080")

}
