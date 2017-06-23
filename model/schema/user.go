package schema

import (
	"github.com/briansan/commune-go/errors"
)

type User struct {
	FirstName string `bson:"fname" json:"fname"`
	LastName  string `bson:"lname" json:"lname"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}

func (u *User) Validate() errors.HTTPError {
	if len(u.Email) == 0 {
		return errors.ValidationErr{Field: "email", Type: "string"}
	}
	if len(u.Username) == 0 {
		return errors.ValidationErr{Field: "username", Type: "string"}
	}
	if len(u.Password) == 0 {
		return errors.ValidationErr{Field: "password", Type: "string"}
	}
	return nil
}
