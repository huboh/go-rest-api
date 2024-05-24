package main

import (
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

// Jwt represents a JWT token as a string.
type Jwt string

// JwtExp represents the token expiration time in Unix format.
type JwtExp int64

// IdToken represents an ID token
type IdToken struct {
	// IdToken represents the jwt token string
	IdToken Jwt `json:"idToken"`

	// IdTokenExpAt represents the token expiration time in Unix format.
	IdTokenExpAt JwtExp `json:"idTokenExpiresAt"`
}

// AuthToken represents an authentication token with its access and refresh tokens
type AuthToken struct {
	// AccessToken represents the jwt token string
	AccessToken Jwt `json:"accessToken"`

	// AccessTokenExpAt represents the token expiration time in Unix format.
	AccessTokenExpAt JwtExp `json:"accessTokenExpiresAt"`

	// RefreshToken represents the jwt token string
	RefreshToken Jwt `json:"refreshToken"`

	// RefreshTokenExpAt represents the token expiration time in Unix format.
	RefreshTokenExpAt JwtExp `json:"refreshTokenExpiresAt"`
}

// TokenConfigs holds the configuration for generating various types of tokens
// including their secrets, issuers, and expiration durations.
type TokenConfigs struct {
	idTokenIssuer    string
	idTokenSecret    []byte
	idTokenExpiresAt time.Duration

	accessTokenIssuer    string
	accessTokenSecret    []byte
	accessTokenExpiresAt time.Duration

	refreshTokenIssuer    string
	refreshTokenSecret    []byte
	refreshTokenExpiresAt time.Duration
}

// NewTokenConfigs initializes a new TokenConfigs instance by reading environment variables
// for issuer, secrets, and expiration durations.
func NewTokenConfigs() *TokenConfigs {
	tc := new(TokenConfigs)
	issuer := os.Getenv("JWT_ISSUER")

	//* set id token configs
	tc.idTokenIssuer = issuer
	tc.idTokenSecret = []byte(os.Getenv("JWT_ID_TOKEN_SECRET"))
	tc.idTokenExpiresAt = Must(time.ParseDuration(os.Getenv("JWT_ID_TOKEN_EXPIRATION")))

	//* set access token configs
	tc.accessTokenIssuer = issuer
	tc.accessTokenSecret = []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
	tc.accessTokenExpiresAt = Must(time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION")))

	//* set refresh tokens configs
	tc.refreshTokenIssuer = issuer
	tc.refreshTokenSecret = []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
	tc.refreshTokenExpiresAt = Must(time.ParseDuration(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION")))

	return tc
}

// createToken generates a Jwt with the provided payload, secret, and expiration duration.
func (tc *TokenConfigs) createToken(p string, s []byte, d time.Duration) (Jwt, JwtExp, error) {
	now := time.Now()
	expAt := now.Add(d)
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    tc.idTokenIssuer,
			Subject:   p,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expAt),
		},
	).SignedString(s)

	if err != nil {
		return "", 0, err
	}

	return Jwt(token), JwtExp(expAt.Unix()), nil
}

// CreateIdToken generates an ID token using the provided payload and the configurations set in TokenConfigs.
func (tc *TokenConfigs) CreateIdToken(payload string) (*IdToken, error) {
	token, expAt, err := tc.createToken(payload, tc.idTokenSecret, tc.idTokenExpiresAt)

	if err != nil {
		return nil, err
	}

	return &IdToken{
		IdToken:      token,
		IdTokenExpAt: expAt,
	}, nil
}

// CreateAuthToken generates an access token and a refresh token concurrently using the provided payload
// and the configurations set in TokenConfigs.
func (tc *TokenConfigs) CreateAuthToken(payload string) (*AuthToken, error) {
	var (
		wg        sync.WaitGroup
		errChan   = make(chan error, 2)
		authToken AuthToken
	)

	wg.Add(2)

	go func() {
		defer wg.Done()

		token, expAt, err := tc.createToken(payload, tc.accessTokenSecret, tc.accessTokenExpiresAt)
		if err != nil {
			errChan <- err
		}

		authToken.AccessToken = token
		authToken.AccessTokenExpAt = expAt
	}()

	go func() {
		defer wg.Done()

		token, expAt, err := tc.createToken(payload, tc.refreshTokenSecret, tc.refreshTokenExpiresAt)
		if err != nil {
			errChan <- err
		}

		authToken.RefreshToken = token
		authToken.RefreshTokenExpAt = expAt
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return &authToken, nil
}
