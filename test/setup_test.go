package test

import (
	"net/http/httptest"
	"os"
	"testing"

	resources "github.com/Pholey/Exgo/resources"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {
	Endpoint string
	Server   *httptest.Server
	Username string
	Email    string
	Name     string
	Password string
}

var _ = Suite(&TestSuite{})

// Setup the server
func (f *TestSuite) SetUpSuite(c *C) {
	// Set our environment to testing
	os.Setenv("DISTRIBUTOR_ENV", "test")

	// Start our server
	conn := httptest.NewServer(resources.NewRouter())

	f.Endpoint = conn.URL
	f.Server = conn
}

// Tear down the server
func (f *TestSuite) TearDownSuite(c *C) {
	f.Server.Close()
}
