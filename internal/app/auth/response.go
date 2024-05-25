package auth

type loginResponse struct {
	Tokens AuthToken `json:"tokens"`
}

type signupResponse struct {
	Tokens AuthToken `json:"tokens"`
}
