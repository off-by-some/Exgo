package user_test

// https://labix.org/gocheck
// https://golang.org/doc/code.html#Testing
import (
    "testing"
    . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type UserTestSuite struct{
  username string
  email string
  name string
  password string
}

var _ = Suite(&UserTestSuite{})

// Setup fixtures
func (s *UserTestSuite) SetUpTest(c *C) {
    s.username = "ReggieSmalls102"
    s.email = "rsmalls@mail.com"
    s.name = "Reginald"
    s.password = "buttcakes123"
}

func (fixture *UserTestSuite) TestHelloWorld(check *C) {
    check.Assert("Reginald", Equals, fixture.name)
}
