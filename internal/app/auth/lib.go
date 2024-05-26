package auth

import (
	"context"
)

var (
	tokens = NewTokenConfigs()
)

func login(ctx context.Context, creds loginCredentials) (loginResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return loginResponse{}, err
	}

	return loginResponse{
		Tokens: *authTokens,
	}, nil
}

func signUp(ctx context.Context) (signupResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return signupResponse{}, err
	}

	return signupResponse{
		Tokens: *authTokens,
	}, nil
}

func refresh(ctx context.Context) (refreshResponse, error) {
	authTokens, err := tokens.CreateAuthToken("userId")

	if err != nil {
		return refreshResponse{}, err
	}

	return refreshResponse{
		Tokens: *authTokens,
	}, nil
}
