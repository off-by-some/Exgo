package session

import (
  "database/sql"
  http "net/http"
  mux "github.com/gorilla/mux"
  sq "github.com/Masterminds/squirrel"
  db "Exgo/db"
  rand "crypto/rand"
  big "math/big"
  pbkdf2 "golang.org/x/crypto/pbkdf2"
  sha256 "crypto/sha256"
)

// FIXME: If the rand stuff in here fails, it
// will probably crash the app, there is no
// error handling here
func hashPass(password string) ([]byte, int, []byte) {
  salt := make([]byte, 32)
  rand.Read(salt)
  ii, _ := rand.Int(rand.Reader, big.NewInt(16000))
  iterations := int(ii.Int64()) + 64000
  hash := pbkdf2.Key([]byte(password), salt, iterations, 32, sha256.New)
  return salt, iterations, hash
}

func createUser(username string, email string, password string, name string) *sql.Rows {
  passwordSalt, passwordIterations, passwordHash := hashPass(password)
  rows, _ := sq.
    Insert("user").
    Columns("username", "email", "password_salt", "password_iterations", "password_hash", "name").
    Values(username, email, passwordSalt, passwordIterations, passwordHash, name).
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
