package test

import . "gopkg.in/check.v1"
import req "github.com/parnurzeal/gorequest"

type User struct {
	Username string
	Email    string
	Name     string
	Password string
}

func (f *TestSuite) TestCreateUser(c *C) {
	u := User{f.Username, f.Email, f.Name, f.Password}
	request := req.New().SetBasicAuth(f.Username, f.Password)
	resp, _, err := request.Post(f.Endpoint + "/user").
		Send(u).
		End()

	if err != nil {
		panic(err)
	}

	c.Assert(resp.StatusCode, Equals, 201)
}
