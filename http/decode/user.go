package decode

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namle133/Log_in2.git/Login_logout/domain"
)

func InputUser(c *gin.Context) *domain.UserInit {
	var user *domain.UserInit
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "400")
		return nil
	}
	if user == nil {
		c.String(http.StatusBadRequest, "400")
		return nil
	}
	return user
}
