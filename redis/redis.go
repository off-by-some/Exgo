package redis

import redis "gopkg.in/redis.v3"

// TODO: Figure out a config setup
var Client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

var response, ConnectionFailed = Client.Ping().Result()
