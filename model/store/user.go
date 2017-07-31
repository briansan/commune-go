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

func newUserQueryByEmail(email string) bson.M {
	return bson.M{"email": email}
}

// GetUsersCollection returns an mgo instance to the users collection
func (m *MongoStore) GetUsersCollection() *mgo.Collection {
	return m.GetDatabase().C(usersCollectionName)
}

// CreateUser inserts user object into db
// error is 500 if mongo fails, 409 if user exists, else nil
func (m *MongoStore) CreateUser(user *schema.User) errors.HTTPError {
	// Check if user exists
	uname := *user.Username
	if user, err := m.GetUserByUsername(uname); err != nil {
		return err
	} else if user != nil {
		return errors.UserExistsErr{uname}
	}

	// Try to insert and return error
	if err := m.GetUsersCollection().Insert(user); err != nil {
		return errors.MongoErr{Err: err}
	}
	return nil
}

// GetUser looks up user in db with given query
func (m *MongoStore) GetUser(q bson.M) (*schema.UserSecure, errors.HTTPError) {
	user := schema.UserSecure{}
	err := m.GetUsersCollection().Find(q).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.MongoErr{Err: err}
	}
	return &user, nil
}

// GetUserByEmail looks up user in db for entire object with given email
// error is 500 if mongo fails, else nil
func (m *MongoStore) GetUserByEmail(email string) (*schema.UserSecure, errors.HTTPError) {
	return m.GetUser(newUserQueryByEmail(email))
}

// GetUserByUsername looks up user in db for entire object with given username
// error is 500 if mongo fails, else nil
func (m *MongoStore) GetUserByUsername(username string) (*schema.UserSecure, errors.HTTPError) {
	return m.GetUser(newUserQueryByUsername(username))
}

// UpdateUser ...
// TODO check for username exists and email exists
// refactor errors to better handle not found case
func (m *MongoStore) UpdateUser(user *schema.User) (*schema.UserSecure, errors.HTTPError) {
	// Do nothing if no username specified
	if user.Username == nil {
		return nil, nil
	}

	// Try to update the user
	q := newUserQueryByUsername(*user.Username)
	changeInfo := mgo.Change{
		Update:    bson.M{"$set": user},
		Upsert:    false,
		ReturnNew: true,
	}
	safeUser := schema.UserSecure{}
	_, err := m.GetUsersCollection().Find(q).Apply(changeInfo, &safeUser)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.MongoErr{Err: err}
	}
	return &safeUser, nil
}

// DeleteUser removes user from db with given username
// error is 500 if mongo fails, else nil
func (m *MongoStore) DeleteUser(username string) (*schema.UserSecure, errors.HTTPError) {
	user, mongoErr := m.GetUserByUsername(username)
	if mongoErr != nil {
		return nil, mongoErr
	}

	if err := m.GetUsersCollection().Remove(newUserQueryByUsername(username)); err != nil {
		return user, errors.MongoErr{Err: err}
	}

	return user, nil
}
