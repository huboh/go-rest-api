package auth

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/huboh/go-rest-api/internal/pkg/env"
	"github.com/huboh/go-rest-api/internal/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// ErrInvalidToken is returned when verifying/validating a Jwt
	ErrInvalidToken = errors.New("invalid token")
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
	issuer := env.Get("JWT_ISSUER")

	//* set id token configs
	tc.idTokenIssuer = issuer
	tc.idTokenSecret = []byte(env.MustGet("JWT_ID_TOKEN_SECRET"))
	tc.idTokenExpiresAt = utils.Must(time.ParseDuration(env.MustGet("JWT_ID_TOKEN_EXPIRATION")))

	//* set access token configs
	tc.accessTokenIssuer = issuer
	tc.accessTokenSecret = []byte(env.MustGet("JWT_ACCESS_TOKEN_SECRET"))
	tc.accessTokenExpiresAt = utils.Must(time.ParseDuration(env.MustGet("JWT_ACCESS_TOKEN_EXPIRATION")))

	//* set refresh tokens configs
	tc.refreshTokenIssuer = issuer
	tc.refreshTokenSecret = []byte(env.MustGet("JWT_REFRESH_TOKEN_SECRET"))
	tc.refreshTokenExpiresAt = utils.Must(time.ParseDuration(env.MustGet("JWT_REFRESH_TOKEN_EXPIRATION")))

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

// verifyToken parses and validates a given token string using the provided secret.
//
// return ErrInvalidToken when t is invalid
func (tc *TokenConfigs) verifyToken(t string, s []byte) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	getSecret := func(token *jwt.Token) (any, error) {
		return s, nil
	}

	token, err := jwt.ParseWithClaims(t, claims, getSecret)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok && !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
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

// VerifyIdToken validates and parses a given token string into an ID token.
//
// return ErrInvalidToken when t is invalid
func (tc *TokenConfigs) VerifyIdToken(token string) (*IdToken, error) {
	claims, err := tc.verifyToken(token, tc.idTokenSecret)
	if err != nil {
		return nil, err
	}

	return &IdToken{
		IdToken:      Jwt(token),
		IdTokenExpAt: JwtExp(claims.ExpiresAt.Unix()),
	}, nil
}

// CreateAuthToken generates an access token and a refresh token concurrently using the provided payload
// and the configurations set in TokenConfigs.
func (tc *TokenConfigs) CreateAuthToken(payload string) (*AuthToken, error) {
	wg := sync.WaitGroup{}
	errChan := make(chan error, 2)
	authToken := AuthToken{}

	wg.Add(2)

	// generate access token
	go func() {
		defer wg.Done()

		token, expAt, err := tc.createToken(payload, tc.accessTokenSecret, tc.accessTokenExpiresAt)
		if err != nil {
			errChan <- err
			return
		}

		authToken.AccessToken = token
		authToken.AccessTokenExpAt = expAt
	}()

	// generate refresh token
	go func() {
		defer wg.Done()

		token, expAt, err := tc.createToken(payload, tc.refreshTokenSecret, tc.refreshTokenExpiresAt)
		if err != nil {
			errChan <- err
			return
		}

		authToken.RefreshToken = token
		authToken.RefreshTokenExpAt = expAt
	}()

	wg.Wait()

	// close so the chan loop exits
	close(errChan)

	// check errors
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return &authToken, nil
}

// VerifyAuthToken validates and parses a given token string into an AuthToken token.
//
// return ErrInvalidToken when t is invalid
func (tc *TokenConfigs) VerifyAuthToken(aToken string, rToken string) (*AuthToken, error) {
	wg := sync.WaitGroup{}
	errChan := make(chan error, 2)
	authToken := AuthToken{}

	wg.Add(2)

	// verify access token
	go func() {
		defer wg.Done()

		claims, err := tc.verifyToken(aToken, tc.accessTokenSecret)

		if err != nil {
			errChan <- err
			return
		}

		// set access token
		authToken.AccessToken = Jwt(aToken)
		authToken.AccessTokenExpAt = JwtExp(claims.ExpiresAt.Unix())
	}()

	// verify refresh token
	go func() {
		defer wg.Done()

		claims, err := tc.verifyToken(rToken, tc.refreshTokenSecret)

		if err != nil {
			errChan <- err
			return
		}

		// set refresh token
		authToken.RefreshToken = Jwt(rToken)
		authToken.RefreshTokenExpAt = JwtExp(claims.ExpiresAt.Unix())
	}()

	wg.Wait()

	// close so the chan loop exits
	close(errChan)

	// check errors
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return &authToken, nil
}
