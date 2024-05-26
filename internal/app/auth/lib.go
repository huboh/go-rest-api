package auth

var (
	tokens = NewTokenConfigs()
)

func login(c loginCredentials) (loginResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		Tokens: *authTokens,
	}, nil
}

func signUp() (signupResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return signupResponse{}, err
	}

	return signupResponse{
		Tokens: *authTokens,
	}, nil
}

func refresh() (refreshResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return refreshResponse{}, err
	}

	return refreshResponse{
		Tokens: *authTokens,
	}, nil
}
