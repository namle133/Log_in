package token

import "github.com/namle133/Log_in2.git/Login_logout/domain"

type Maker interface {
	CreateToken(u *domain.UserInit) (string, *Payload, error)
}
