package v1

type User struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

func CreateUser(u *User) error {
	return nil
}
