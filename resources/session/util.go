package session

import (
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

func verifyHash(password string, salt []byte, iterations int, hash []byte) bool {
  compareHash := pbkdf2.Key([]byte(password), salt, iterations, 32, sha256.New)
  return bytes.Equal(hash, compareHash)
}
