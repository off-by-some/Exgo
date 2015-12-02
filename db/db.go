package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  ssq "github.com/Masterminds/squirrel"
)

var Client, DatabaseConnFailed = sql.Open("postgres", "postgres://postgres:@localhost/exgo?sslmode=disable")
// Set up squirrel with statement caching and dollar sign placeholders
var Sq = ssq.StatementBuilder.RunWith(ssq.NewStmtCacher(Client)).PlaceholderFormat(ssq.Dollar)
