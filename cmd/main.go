package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/namle133/Log_in2.git/Login_logout/database"
	"github.com/namle133/Log_in2.git/Login_logout/http/decode"
	"github.com/namle133/Log_in2.git/Login_logout/http/encode"
	"github.com/namle133/Log_in2.git/Login_logout/service"
)

func main() {
	r := gin.Default()

	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	us := &service.UserService{Db: database.ConnectDatabase(host, user, pw, name, port)}
	var i service.IUser = us
	err := us.UserAdmin()
	if err != nil {
		fmt.Errorf("can't create useradmin with err: %v", err)
		return
	}

	r.POST("/signin", func(c *gin.Context) {
		u := decode.InputUser(c)
		payload, err := i.SignIn(c, u)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
			return
		}
		encode.SignInResponse(c, payload)
	})

	r.POST("/createuser", func(c *gin.Context) {
		//Authorization Bearer Token
		token := c.GetHeader("Authorization")[7:]
		err := i.CheckUserAdmin(c, token)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
			return
		}
		u := decode.InputUser(c)
		er := i.CreateUser(c, u)
		if er != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
			return
		}
		encode.CreateUserResponse(c)
	})

	r.DELETE("/logout", func(c *gin.Context) {
		//Authorization Bearer Token
		token := c.GetHeader("Authorization")[7:]
		err := i.LogOut(c, token)
		if err != nil {
			c.String(http.StatusBadRequest, "LogOut Failed")
			return
		}
		encode.LogoutResponse(c)
	})
	r.Run(":8000")
}
