package datastore

import (
	"github.com/shunsukw/go-chat/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(dbtype int, dbConnectionString string) (Datastore, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDatastore(dbConnectionString)

	}

	return nil, nil
}
