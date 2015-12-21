package test

// http://dennissuratna.com/testing-in-go/
// https://labix.org/gocheck
// https://golang.org/doc/code.html#Testing
import (
    "testing"
    "fmt"
    . "gopkg.in/check.v1"
    "net/http/httptest"
    resources "Exgo/resources"
    "os"
    req "github.com/parnurzeal/gorequest"
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

func TestMain(m *testing.M) {
  conn := httptest.NewServer(resources.NewRouter())
  fmt.Printf("Server set")

  // Close the socket after we are done testing
  val := m.Run()

  os.Exit(val)
  conn.Close()
}

// Setup fixtures
func (s *UserTestSuite) SetUpTest(c *C) {

}

func (fixture *UserTestSuite) TestCreateUser(c *C) {
  request := req.New().SetBasicAuth(fixture.username, fixture.password)
  resp, _, err := request.Post("http://0.0.0.0:8080/user").
    Send(fixture).
    End()

  if (err != nil) {
    panic(err)
  }

  c.Assert(resp.StatusCode, Equals, 201)
}
