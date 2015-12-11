package db

import (
  "database/sql"
  "Exgo/config"
  "fmt"
  _ "github.com/lib/pq"
  ssq "github.com/Masterminds/squirrel"
)

type DatabaseConnection struct{
  user string
  host string
  database string
}

var Client *sql.DB
var Sq ssq.StatementBuilderType
var ConnectionInfo DatabaseConnection

func init() {
  info := config.File.GetStringMap("database")

  ConnectionInfo = DatabaseConnection{
    info["user"].(string),
    info["host"].(string),
    info["database"].(string),
  }

  connString := fmt.Sprintf(
    "postgres://%s:@%s/%s?sslmode=disable",
    info["user"],
    info["host"],
    info["database"],
  )

  client, err := sql.Open("postgres", connString)

  if (err != nil) {
    panic(err)
  }

  Client = client
  println("Connected to Postgres")
}
