package service

import (
	"context"

	"github.com/namle133/Log_in2.git/Login_logout/domain"
	"github.com/namle133/Log_in2.git/Login_logout/token"
)

type IUser interface {
	CreateUser(c context.Context, u *domain.UserInit) error
	SignIn(c context.Context, u *domain.UserInit) (*token.Payload, string, error)
	UserAdmin() error
	CheckUserAdmin(c context.Context, token string, username string) error
	LogOut(c context.Context, token string, username string) error
}
