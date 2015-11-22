package session

import (
  "database/sql"
  http "net/http"
  sq "github.com/Masterminds/squirrel"
  db "Exgo/db"
)

func createUser(username string, email string, password string, name string) *sql.Rows {
  rows, _ := sq.
    Insert("user").
    Columns("username", "email", "password", "name").
    Values(username, email, password, name).
    RunWith(db.Client).
    Query()

  defer rows.Close()

  return rows
}

func Create(res http.ResponseWriter, req *http.Request) {
  createUser("rsmalls22", "reggie@mail.com", "suchsecretmuchsecure", "Reginald Smalls")
  // createUser(vars["username"], vars["email"], vars["password"], vars["name"])


}
