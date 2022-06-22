package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/namle133/Log_in2.git/Login_logout/domain"
)

var JwtKey = []byte("my-secrect-key")

type Payload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JwtMaker struct{}

func (m *JwtMaker) CreateToken(u *domain.UserInit) (string, *Payload, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	p := &Payload{
		Username: u.Username,
		Email:    u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", nil, err
	}
	return tokenString, p, nil
}

func (m *JwtMaker) CheckTokenValid(tknStr string) (*Payload, error) {
	payload := &Payload{}
	tkn, err := jwt.ParseWithClaims(tknStr, payload, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, fmt.Errorf("Token invalid")
	}
	return payload, nil
}
