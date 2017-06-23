package store

import (
	"fmt"

	"github.com/mgutz/logxi/v1"
	"gopkg.in/mgo.v2"
)

const (
	envMongoAuth = "COMMUNE_MONGO_AUTH"
	envMongoHost = "COMMUNE_MONGO_HOST"
)

var (
	logger = log.New("store")

	mongoAuth    = "mdbcommune:commune"
	mongoHost    = "localhost:27017"
	databaseName = "commune"

	mongo *mgo.Session
)

type MongoStore struct {
	s *mgo.Session
}

func getMongoURL() string {
	return fmt.Sprintf("mongodb://%v@%v/%v", mongoAuth, mongoHost, databaseName)
}

func init() {
	var err error
	if mongo, err = mgo.Dial(getMongoURL()); err != nil {
		panic(err)
	}
}

func NewMongoStore() *MongoStore {
	return &MongoStore{mongo.Copy()}
}

func (m *MongoStore) Cleanup() {
	m.s.Close()
}

func (m *MongoStore) GetDatabase() *mgo.Database {
	return m.s.DB("commune")
}
