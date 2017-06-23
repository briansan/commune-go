package store

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/briansan/commune-go/model/schema"
)

type StoreTestSuite struct {
	suite.Suite
	mongo *MongoStore
}

func (suite *StoreTestSuite) SetupTest() {
	databaseName = "test"

	var err error
	suite.mongo, err = NewMongoStore()
	suite.Nil(err)
}

func (suite *StoreTestSuite) TearDownTest() {
	suite.mongo.GetUsersCollection().DropCollection()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

func (suite *StoreTestSuite) Test001_User() {
	username := "foo"

	err := suite.mongo.CreateUser(&schema.User{Username: username})
	suite.Nil(err)

	exists, err := suite.mongo.UserExists(username)
	suite.Nil(err)
	suite.True(exists)

	user, err := suite.mongo.GetUserByUsername(username)
	suite.Nil(err)
	suite.Equal(username, user.Username)

	sameUser, err := suite.mongo.DeleteUser(username)
	suite.Nil(err)
	suite.Equal(*user, *sameUser)
}
