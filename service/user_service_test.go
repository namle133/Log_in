package service

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/namle133/Log_in2.git/Login_logout/database"
	"github.com/namle133/Log_in2.git/Login_logout/domain"
	"github.com/namle133/Log_in2.git/Login_logout/token"
	"gorm.io/gorm"
)

var (
	Host = os.Getenv("DB_HOST")
	User = os.Getenv("DB_USER_TEST")
	Pw   = os.Getenv("DB_PASSWORD_TEST")
	Name = os.Getenv("DB_NAM_TEST")
	Port = os.Getenv("DB_PORT_TEST")
)

func TestUserService_SignIn(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}

	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c  context.Context
		ui *domain.UserInit
	}
	tests := []struct {
		fields  fields
		args    args
		want    *token.Payload
		wantErr bool
	}{
		//case success
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "admin",
					Password: "admin1234",
					Email:    "admin@gmail.com"},
			},
			want: &token.Payload{
				Username:       "admin",
				Email:          "admin@gmail.com",
				StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(5 * time.Minute).Unix()},
			},
			wantErr: false,
		},

		//case failed
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "",
					Password: "Namle311",
					Email:    "Namle@gmail.com"},
			},
			want: &token.Payload{
				Username:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "",
					Password: "Namle311",
					Email:    "Namle@gmail.com"},
			},
			want: &token.Payload{
				Username:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "",
					Password: "",
					Email:    ""},
			},
			want: &token.Payload{
				Username:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "Namle",
					Password: "",
					Email:    "Namle@gmail.com"},
			},
			want: &token.Payload{
				Username:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				ui: &domain.UserInit{
					Username: "Namle",
					Password: "Namle311",
					Email:    ""},
			},
			want: &token.Payload{
				Username:       "",
				Email:          "",
				StandardClaims: jwt.StandardClaims{ExpiresAt: 0},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		_, _, err := us.SignIn(tt.args.c, tt.args.ui)
		if (err != nil) != tt.wantErr {
			t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func TestUserService_CreateUser(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c context.Context
		u *domain.UserInit
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		//case success
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "Namle",
					Password: "Namle1234",
					Email:    "Namle@gmail.com"},
			},
			wantErr: false,
		},

		//case failed
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "",
					Password: "31231423",
					Email:    ""},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "",
					Password: "31231423",
					Email:    "nam@gmail.com"},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "Namle",
					Password: "",
					Email:    ""},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "Nam",
					Password: "",
					Email:    "nam@gmail.com"},
			},
			wantErr: true,
		},

		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c: context.Background(),
				u: &domain.UserInit{
					Username: "",
					Password: "",
					Email:    ""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.CreateUser(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
			t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_UserAdmin(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}
	type fields struct {
		Db *gorm.DB
	}
	tests := []struct {
		fields  fields
		wantErr bool
	}{
		{
			fields:  fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.UserAdmin(); (err != nil) != tt.wantErr {
			t.Errorf("UserAdmin() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_CheckUserAdmin(t *testing.T) {
	var g *gin.Context
	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c        context.Context
		token    string
		username string
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c:     context.Background(),
				token: g.GetHeader("Authorization"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.CheckUserAdmin(tt.args.c, tt.args.token, tt.args.username); (err != nil) != tt.wantErr {
			t.Errorf("CheckUserAdmin() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestUserService_LogOut(t *testing.T) {
	var g *gin.Context
	e := godotenv.Load()
	if e != nil {
		log.Fatal("error loading .env file")
	}
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		c        context.Context
		token    string
		username string
	}
	tests := []struct {
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{Db: database.ConnectDatabase(Host, User, Pw, Name, Port)},
			args: args{
				c:     context.Background(),
				token: g.GetHeader("Authorization"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		us := &UserService{
			Db: tt.fields.Db,
		}
		if err := us.LogOut(tt.args.c, tt.args.token, tt.args.username); (err != nil) != tt.wantErr {
			t.Errorf("LogOut() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
