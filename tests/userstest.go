package tests

import (
	"github.com/revel/revel/testing"
)

type UsersTest struct {
	testing.TestSuite
}

func (t *UsersTest) Before() {
	// TODO: Setup test DB - seed appropiately.
	println("Set up")
}

func (t *UsersTest) After() {
	// TODO: Tear down (https://github.com/jinzhu/gorm/issues/45)
	println("Tear down")
}

func (t *UsersTest) TestUsersIndex() {
	t.Get("/users")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}