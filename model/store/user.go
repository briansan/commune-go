package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/briansan/commune-go/errors"
	"github.com/briansan/commune-go/model/schema"
)

const (
	usersCollectionName = "users"
)

func newUserQueryByUsername(username string) bson.M {
	return bson.M{"username": username}
}

func (m *MongoStore) GetUsersCollection() *mgo.Collection {
	return m.GetDatabase().C(usersCollectionName)
}

func (m *MongoStore) UserExists(username string) (bool, error) {
	coll := m.GetUsersCollection()
	n, err := coll.Find(newUserQueryByUsername(username)).Count()
	if err != nil {
		return false, err
	}

	return n == 1, nil
}

func (m *MongoStore) CreateUser(user *schema.User) errors.HTTPError {
	// Check if user exists
	uname := user.Username
	if exists, err := m.UserExists(uname); err != nil {
		return errors.MongoErr{Err: err}
	} else if exists {
		return errors.UserExistsErr{uname}
	}

	// Try to insert and return error
	if err := m.GetUsersCollection().Insert(user); err != nil {
		return errors.MongoErr{Err: err}
	}
	return nil
}

func (m *MongoStore) GetUserByUsername(username string) (*schema.User, error) {
	user := schema.User{}
	err := m.GetUsersCollection().Find(newUserQueryByUsername(username)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoStore) DeleteUser(username string) (*schema.User, error) {
	user, err := m.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = m.GetUsersCollection().Remove(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
