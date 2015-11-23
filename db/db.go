package db

import (
  "database/sql"
  _ "github.com/lib/pq"
)

var Client, DatabaseConnFailed = sql.Open("postgres", "postgres://postgres:@localhost?sslmode=disable")
