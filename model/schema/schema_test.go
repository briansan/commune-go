package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test001_UserValidation(t *testing.T) {
	var err error
	u := User{}

	// Test email field
	err = u.Validate()
	assert.Equal(t, "email field is required as string", err.Error())

	// Test username field
	u.Email = "foo"
	err = u.Validate()
	assert.Equal(t, "username field is required as string", err.Error())

	// Test password field
	u.Username = "bar"
	err = u.Validate()
	assert.Equal(t, "password field is required as string", err.Error())

	// Test good
	u.Password = "baz"
	err = u.Validate()
	assert.Nil(t, err)
}
