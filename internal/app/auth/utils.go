package auth

import (
	"fmt"
	"net/http"
	"regexp"
)

var tknRegexp = regexp.MustCompile("^Bearer\x20(.+)$")

func getAuthHeaderToken(r http.Request) (string, error) {
	matches := tknRegexp.FindStringSubmatch(r.Header.Get("Authorization"))

	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("malformed authorization header")
	}

	return matches[1], nil
}
