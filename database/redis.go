package database

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/namle133/Log_in2.git/Login_logout/token"
)

func NewConn(host string, port string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})

	return rdb
}

func Set(client *redis.Client, t *token.Payload, tknStr string) error {

	// Set Time-To-Live: 5* time.Minute
	if err := client.Set(tknStr, t.Username, 5*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func Get(client *redis.Client, tknStr string) (string, error) {

	value, err := client.Get(tknStr).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func Delete(client *redis.Client, tknStr string) error {
	err := client.Del(tknStr).Err()
	if err != nil {
		return err
	}
	return nil

}
