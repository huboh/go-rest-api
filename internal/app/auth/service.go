package auth

func login() (loginResponse, error) {
	tokens := NewTokenConfigs()
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		Tokens: *authTokens,
	}, nil
}

func signUp() (signupResponse, error) {
	tokens := NewTokenConfigs()
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return signupResponse{}, err
	}

	return signupResponse{
		Tokens: *authTokens,
	}, nil
}

func refresh() (refreshResponse, error) {
	tokens := NewTokenConfigs()
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return refreshResponse{}, err
	}

	return refreshResponse{
		Tokens: *authTokens,
	}, nil
}
