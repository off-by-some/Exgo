package session

import (
  "database/sql"
  http "net/http"
  mux "github.com/gorilla/mux"
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

func create(res http.ResponseWriter, req *http.Request) {
  // TODO: Validation
  vars := mux.Vars(req)
  // createUser("rsmalls22", "reggie@mail.com", "suchsecretmuchsecure", "Reginald Smalls")
  createUser(vars["username"], vars["email"], vars["password"], vars["name"])
}

func HandleLogin(res http.ResponseWriter, req *http.Request) {


}
