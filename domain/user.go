package domain

type UserInit struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
	Email    string `json:"email"`
}

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"tokenString"`
}
