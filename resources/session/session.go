package session

import (
	json "encoding/json"
	"io"
	"io/ioutil"
	http "net/http"

	db "github.com/Pholey/Exgo/db"

	sq "github.com/Masterminds/squirrel"

	"github.com/gin-gonic/gin"
)

// User - Struct for dealing with users
type User struct {
	Salt       []byte `db:"salt"`
	Iterations int    `db:"password_iterations"`
	Hash       int    `db:"password_hash"`
	Password   string `json:"password"`
	UserName   string `db:"username" json:"userName"`
	Name       string `db:"name"     json:"name"`
	Email      string `db:"email"    json:"email"`
}

// Users - Our array of users
type Users []User

// Create - Create a user
func Create(c *gin.Context) {
	// TODO(pholey): Abstract this out or find a better lib
	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))

	if err != nil {
		// TODO(pholey): Proper error handing
		panic(err)
	}

	if err := c.Request.Body.Close(); err != nil {
		panic(err)
	}

	// Grab our user data
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		panic(err)
	}

	// TODO(pholey): Password validation, refactoring
	// Hash our password
	salt, iterations, hash := hashPass(user.Password)

	// Create our statement
	sql, args, _ := db.Sq.
		Insert("\"user\"").
		Columns("username", "name", "email", "password_hash",
		"password_iterations", "password_salt").
		Values(user.UserName, user.Name, user.Email, hash, iterations, salt).
		ToSql()

	rows, err := db.Client.Query(sql, args...)

	_ = (rows)

	if err != nil {
		panic(err)
	}

	c.Writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	c.Writer.WriteHeader(http.StatusCreated)
}

// TODO: Handle non-existant users
func getUserAuthInfo(username string) ([]byte, int, []byte) {
	var (
		salt       []byte
		iterations int
		hash       []byte
	)
	rows, _ := db.Sq.
		Select("password_salt", "password_iterations", "password_hash").
		From("\"user\"").
		Where(sq.Eq{"username": username}).
		Query()
	rows.Scan(&salt, &iterations, &hash)
	return salt, iterations, hash
}

func auth(username string, password string) bool {
	salt, iterations, hash := getUserAuthInfo(username)
	return verifyHash(password, salt, iterations, hash)
}
