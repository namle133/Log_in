package service

import (
	"context"
	"fmt"

	"github.com/namle133/Log_in2.git/Login_logout/domain"
	"github.com/namle133/Log_in2.git/Login_logout/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func hash(s string) []byte {
	bsp, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return bsp
}

func ComparePassword(hashedPassword []byte, password []byte) error {
	er := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if er != nil {
		return er
	}
	return nil
}

func (us *UserService) SignIn(c context.Context, ui *domain.UserInit) (*token.Payload, error) {
	if ui == nil {
		return nil, ErrNotFound
	}
	if ui.Username == "" {
		return nil, ErrNameIsRequired
	}
	if ui.Email == "" {
		return nil, ErrEmailIsRequired
	}
	if ui.Password == "" {
		return nil, ErrPasswordIsRequired
	}
	var u *domain.User
	e := us.Db.First(&u, "username=? and email=?", ui.Username, ui.Email).Scan(&u).Error
	if e != nil {
		return nil, e
	}

	err := ComparePassword(u.Password, []byte(ui.Password))
	if err != nil {
		return nil, err
	} else {
		fmt.Println("password are equal")
	}
	var m token.Maker = &token.JwtMaker{}
	tokenString, Payload, err := m.CreateToken(ui)
	if err != nil {
		return nil, err
	}

	tk := &domain.Token{Username: ui.Username, TokenString: tokenString}
	failed := us.Db.Create(tk).Error
	if failed != nil {
		return nil, failed
	}
	return Payload, nil
}

func (us *UserService) CreateUser(c context.Context, ui *domain.UserInit) error {
	var u *domain.User
	if ui == nil {
		return ErrNotFound
	}
	if ui.Username == "" {
		return ErrNameIsRequired
	}
	if ui.Email == "" {
		return ErrEmailIsRequired
	}
	if ui.Password == "" {
		return ErrPasswordIsRequired
	}
	e := us.Db.First(&u, "username=? or email=?", ui.Username, ui.Email).Scan(&u).Error
	if e == nil {
		return ErrUserIsExist
	}
	uh := &domain.User{Username: ui.Username, Password: hash(ui.Password), Email: ui.Email}
	err := us.Db.Create(uh).Error
	if err == nil {
		return err
	}
	return nil
}

func (us *UserService) UserAdmin() error {
	u := domain.UserInit{Username: "admin", Password: "admin1234", Email: "admin@gmail.com"}
	uh := &domain.User{Username: u.Username, Password: hash(u.Password), Email: u.Email}
	err := us.Db.Create(uh).Error
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) CheckUserAdmin(c context.Context, tknStr string) error {
	var t *domain.Token
	var m token.Maker = &token.JwtMaker{}
	er := m.CheckTokenValid(tknStr)
	if er != nil {
		return ErrTokenIsInvalid
	}
	err := us.Db.First(&t, "token_string = ?", tknStr).Error
	if err != nil {
		return err
	}
	if t.Username != "admin" {
		return ErrRecordNotFound
	}
	return nil
}

func (us *UserService) LogOut(c context.Context, token string) error {
	var t *domain.Token
	if token == "" {
		return ErrNotFound
	}
	err := us.Db.Where("token_string = ?", token).Delete(&t).Error
	if err != nil {
		return err
	}
	return nil
}
