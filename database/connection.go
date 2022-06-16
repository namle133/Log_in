package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(host string, user string, pw string, name string, port string) *gorm.DB {
	dsn := "host=" + host + " user=" + user + " password=" + pw + " dbname=" + name + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().CurrentDatabase()
	return db

}
