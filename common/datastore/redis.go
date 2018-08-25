package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/shunsukw/go-chat/models"
)

// RedisDatastore ...
type RedisDatastore struct {
	*pool.Pool
}

// NewRedisDatastore ...
func NewRedisDatastore(address string) (*RedisDatastore, error) {
	connectionPool, err := pool.New("tcp", address, 10)
	if err != nil {
		return nil, err
	}

	return &RedisDatastore{
		Pool: connectionPool,
	}, nil
}

// CreateUser ...
func (r *RedisDatastore) CreateUser(user *models.User) error {
	userJSON, err := json.Marshal(*user)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "user:"+user.Username, string(userJSON)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

// GetUser ...
func (r *RedisDatastore) GetUser(username string) (*models.User, error) {

	exists, err := r.Cmd("EXISTS", "user:"+username).Int()

	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	var u models.User

	userJSON, err := r.Cmd("GET", "user:"+username).Str()

	fmt.Println("userJSON: ", userJSON)

	if err != nil {
		log.Print(err)

		return nil, err
	}

	if err := json.Unmarshal([]byte(userJSON), &u); err != nil {
		log.Print(err)
		return nil, err
	}

	return &u, nil
}

// Close ...
func (r *RedisDatastore) Close() {

	r.Close()
}
