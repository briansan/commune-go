package schema

import (
	"github.com/briansan/commune-go/errors"
)

type UserSecure struct {
	FirstName string `bson:"fname" json:"fname"`
	LastName  string `bson:"lname" json:"lname"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
	Username  string `bson:"username" json:"username"`
}

type User struct {
	FirstName *string `bson:"fname,omitempty" json:"fname,omitempty"`
	LastName  *string `bson:"lname,omitempty" json:"lname,omitempty"`
	Email     *string `bson:"email,omitempty" json:"email,omitempty"`
	Phone     *string `bson:"phone,omitempty" json:"phone,omitempty"`
	Username  *string `bson:"username,omitempty" json:"username,omitempty"`
	Password  *string `bson:"password,omitempty" json:"password,omitempty"`
}

func (u *User) Validate() errors.HTTPError {
	if u.Email == nil || len(*u.Email) == 0 {
		return errors.ValidationErr{Field: "email", Type: "string"}
	}
	if u.Username == nil || len(*u.Username) == 0 {
		return errors.ValidationErr{Field: "username", Type: "string"}
	}
	if u.Password == nil || len(*u.Password) == 0 {
		return errors.ValidationErr{Field: "password", Type: "string"}
	}
	return nil
}
