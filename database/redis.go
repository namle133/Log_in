package database

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/namle133/Log_in2.git/Login_logout/token"
)

func NewConn(host string, port string, pw string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pw,
		DB:       0,
	})

	return rdb
}

func Set(client *redis.Client, t *token.Payload, tknStr string) error {

	// Set Time-To-Live: 5* time.Minute
	if err := client.Set(t.Username, tknStr, 5*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func Get(client *redis.Client, username string) (string, error) {

	value, err := client.Get(username).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func Delete(client *redis.Client, username string) error {
	err := client.Del(username).Err()
	if err != nil {
		return err
	}
	return nil

}
