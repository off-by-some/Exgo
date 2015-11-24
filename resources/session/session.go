package session

import (
  "database/sql"
  http "net/http"
  sq "github.com/Masterminds/squirrel"
  db "Exgo/db"
  rand "crypto/rand"
  big "math/big"
  pbkdf2 "golang.org/x/crypto/pbkdf2"
  sha256 "crypto/sha256"
  "bytes"
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

func Create(res http.ResponseWriter, req *http.Request) {
  createUser("rsmalls22", "reggie@mail.com", "suchsecretmuchsecure", "Reginald Smalls")
  // createUser(vars["username"], vars["email"], vars["password"], vars["name"])
}

func verifyHash(password string, salt []byte, iterations int, hash []byte) bool {
  compareHash := pbkdf2.Key([]byte(password), salt, iterations, 32, sha256.New)
  return bytes.Equal(hash, compareHash)
}

// I'm assuming usernames will be unique
// Really there should be a unique constraint
// in the db on something we can use to grab
// a user
func getUser(username string) (string, string, string, []byte, int, []byte, string) {
  var (
    id string
    un string
    email string
    salt []byte
    iterations int
    hash []byte
    body string
  )

  rows, _ := sq.
    Select("*").
    From("user").
    Where(sq.Eq{"username": username}).
    RunWith(db.Client).
    Query()
  rows.Scan(&id, &un, &email, &salt, &iterations, &hash, &body)
  return id, un, email, salt, int(iterations), hash, body
}

func auth(username string, password string) bool {
  id, un, email, salt, iterations, hash, body := getUser(username)
  return verifyHash(password, salt, iterations, hash)
}
