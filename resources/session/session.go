package session

import (
	redis "Exgo/redis"
	json "encoding/json"
	"io"
	"io/ioutil"
	http "net/http"

	sq "github.com/Masterminds/squirrel"
	db "github.com/Pholey/Exgo/db"
)

type User struct {
	Salt       []byte `db:"salt"`
	Iterations int    `db:"password_iterations"`
	Hash       int    `db:"password_hash"`
	Password   string `json:"password"`
	UserName   string `db:"username" json:"userName"`
	Name       string `db:"name"     json:"name"`
	Email      string `db:"email"    json:"email"`
}

type Users []User

func unmarshalFromRequest(req *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		// TODO(pholey): Proper error handing
		panic(err)
	}

	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	// Grab our user data
	var user User
	if err := json.Unmarshal(body, &v); err != nil {
		panic(err)
	}
}

func Create(res http.ResponseWriter, req *http.Request) {
	var user User
	unmarshalFromRequest(req, &user)

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

	res.WriteHeader(http.StatusCreated)
}

// TODO: Handle non-existant users
func getUserAuthInfo(username string) (string, []byte, int, []byte) {
	var (
		id         string
		salt       []byte
		iterations int
		hash       []byte
	)
	rows, _ := db.Sq.
		Select("id", "password_salt", "password_iterations", "password_hash").
		From("\"user\"").
		Where(sq.Eq{"username": username}).
		QueryRow()
	rows.Scan(&id, &salt, &iterations, &hash)
	return id, salt, iterations, hash
}

type Token struct {
	token string
}

// func auth(username string, password string) bool {
func Auth(res http.ResponseWriter, req *http.Request) {
	var user User
	unmarshalFromRequest(req, &user)

	id, salt, iterations, hash := getUserAuthInfo(user.UserName)
	if verifyHash(password, salt, iterations, hash) {
		token := Token(randBase64String(16))
		redis.Client.Set(token.token, id, 0)

		res.Header().Set("Content-Type", "application/json;charset=UTF-8")
		res.WriteHeader(http.StatusCreated)

		js, _ := json.Marshal(token)
		res.Write(js)
	}
}

func GetUserFromToken(token) string {
	return redis.Client.get(token)
}
