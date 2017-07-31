package store

import (
	"fmt"

	"github.com/mgutz/logxi/v1"
	"gopkg.in/mgo.v2"

	"github.com/briansan/commune-go/errors"
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
	if err := InitMongoSession(); err != nil {
		panic(err)
	}
}

// InitMongoSession resets the mongo session pointer with updated connection info
func InitMongoSession() error {
	// To avoid a socket leak
	if mongo != nil {
		CleanupMongoSession()
	}

	// Establish new session
	var err error
	if mongo, err = mgo.Dial(getMongoURL()); err != nil {
		return err
	}
	return nil
}

// CleanupMongoSession closes the current session and sets the pointer to nil
func CleanupMongoSession() {
	if mongo == nil {
		return
	}
	mongo.Close()
	mongo = nil
}

// NewMongoStore returns an instance of the store with a copied mongo session
// error is 500 if mongo ping fails
func NewMongoStore() (*MongoStore, errors.HTTPError) {
	if err := mongo.Ping(); err != nil {
		return nil, errors.MongoErr{Err: err}
	}
	return &MongoStore{mongo.Copy()}, nil
}

// Cleanup closes the mongo session of this store object
func (m *MongoStore) Cleanup() {
	m.s.Close()
}

// GetDatabase returns a pointer to an mgo database object
func (m *MongoStore) GetDatabase() *mgo.Database {
	return m.s.DB(databaseName)
}
