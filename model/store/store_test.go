package store

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/briansan/commune-go/errors"
	"github.com/briansan/commune-go/model/schema"
)

type StoreTestSuite struct {
	suite.Suite
	store *MongoStore
}

func (suite *StoreTestSuite) SetupTest() {
	// Use test database and reestablish session
	databaseName = "commune_test"
	InitMongoSession()

	var err errors.HTTPError
	suite.store, err = NewMongoStore()
	suite.Nil(err)

	suite.store.GetUsersCollection().DropCollection()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

// Test001_User asserts proper CRUD functionality of user object with mongo
func (suite *StoreTestSuite) Test001_User() {
	username := "foo"
	email := "bar"

	// Test CreateUser
	err := suite.store.CreateUser(&schema.User{
		Username: &username,
		Email:    &email,
	})
	suite.Nil(err)

	// Test GetUserByUsername
	user, err := suite.store.GetUserByUsername(username)
	suite.Nil(err)
	suite.NotNil(user)
	suite.Equal(username, user.Username)
	suite.Equal(email, user.Email)

	// Test GetUserByEmail
	user, err = suite.store.GetUserByEmail(email)
	suite.Nil(err)
	suite.NotNil(user)
	suite.Equal(username, user.Username)
	suite.Equal(email, user.Email)

	// Test UpdateUser
	firstName := "fname"
	userPatch := schema.User{Username: &username, FirstName: &firstName}
	user, err = suite.store.UpdateUser(&userPatch)
	suite.Nil(err)
	suite.Equal(user.Username, *userPatch.Username)
	suite.Equal(user.FirstName, *userPatch.FirstName)
	suite.Equal(user.Email, email)

	// Test DeleteUser
	user, err = suite.store.DeleteUser(username)
	suite.Nil(err)
	suite.NotNil(user)
	suite.Equal(username, user.Username)
	suite.Equal(email, user.Email)
}
