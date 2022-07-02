package database

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/namle133/Log_in2.git/Login_logout/token"
)

type redisClient struct {
	c *redis.Client
}

var Cl = &redisClient{}

func NewConn(host string, port string) (*redisClient, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Ping().Err()
	if err != nil {
		return nil, err
	}
	Cl.c = rdb

	return Cl, nil
}

func (Cl *redisClient) Set(t *token.Payload, tknStr string) error {
	// Set Time-To-Live: 5* time.Minute
	fmt.Println(t.Username, tknStr)
	if err := Cl.c.Set(t.Username, tknStr, 5*time.Minute).Err(); err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}

func (Cl *redisClient) Get(username string) (string, error) {

	value, err := Cl.c.Get(username).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (Cl *redisClient) Delete(username string) error {
	err := Cl.c.Del(username).Err()
	if err != nil {
		return err
	}
	return nil

}
