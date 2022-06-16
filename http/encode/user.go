package encode

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namle133/Log_in2.git/Login_logout/token"
)

func SignInResponse(c *gin.Context, claims *token.Payload) {
	c.String(http.StatusOK, fmt.Sprintf("Welcome to %v", claims.Username))
}

func CreateUserResponse(c *gin.Context) {
	c.String(http.StatusOK, "SignUp Success")
}

func LogoutResponse(c *gin.Context) {
	c.String(http.StatusOK, "Old cookie deleted. Logged out!")
}
