package redis

import redis "gopkg.in/redis.v3"
import "github.com/Pholey/Exgo/config"
import "fmt"

var Client *redis.Client
func init() {
  info := config.File.GetStringMap("redis")
  port := config.File.GetInt("redis.port")
  database := config.File.GetInt("redis.database")
  Client = redis.NewClient(&redis.Options{
      Addr:     fmt.Sprintf("%s:%d", info["host"].(string), port),
      Password: "",
      DB:       int64(database),
  })

  var _, err = Client.Ping().Result()

  if (err != nil) {
    panic(err)
  }

  println("Connected to Redis at port " + fmt.Sprintf("%d", port))
}
