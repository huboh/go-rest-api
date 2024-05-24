package main

type AuthResponse struct {
	Tokens AuthToken `json:"tokens"`
}

func Login() (AuthResponse, error) {
	tokens := NewTokenConfigs()
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		Tokens: *authTokens,
	}, nil
}

func SignUp() (AuthResponse, error) {
	tokens := NewTokenConfigs()
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		Tokens: *authTokens,
	}, nil
}
