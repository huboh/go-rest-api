package auth

type loginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
