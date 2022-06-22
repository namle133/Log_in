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
	"github.com/namle133/Log_in2.git/Login_logout/token"
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
	re_host := os.Getenv("REDIS_HOST")
	re_port := os.Getenv("REDIS_PORT")
	re_pw := os.Getenv("REDIS_PASSWORDS")

	// New connect
	client := database.NewConn(re_host, re_port, re_pw)
	us := &service.UserService{Db: database.ConnectDatabase(host, user, pw, name, port)}

	var i service.IUser = us
	err := us.UserAdmin()
	if err != nil {
		fmt.Errorf("can't create useradmin with err: %v", err)
		return
	}

	r.POST("/signin", func(c *gin.Context) {
		u := decode.InputUser(c)
		payload, tknStr, err := i.SignIn(c, u)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
			return
		}

		database.Set(client, payload, tknStr)
		encode.SignInResponse(c, payload)
	})

	r.POST("/createuser", func(c *gin.Context) {
		//Authorization Bearer Token
		tknStr := c.GetHeader("Authorization")[7:]

		var m token.Maker = &token.JwtMaker{}
		payload, e := m.CheckTokenValid(tknStr)
		if e != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", e))
			return
		}

		err := i.CheckUserAdmin(c, tknStr, payload.Username)
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
		tknStr := c.GetHeader("Authorization")[7:]

		var m token.Maker = &token.JwtMaker{}
		payload, e := m.CheckTokenValid(tknStr)
		if e != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", e))
			return
		}

		err := i.LogOut(c, tknStr, payload.Username)
		if err != nil {
			c.String(http.StatusBadRequest, "LogOut Failed")
			return
		}

		er := database.Delete(client, payload.Username)
		if er != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%v", er))
			return
		}
		encode.LogoutResponse(c)
	})
	r.Run(":8000")
}
