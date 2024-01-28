// auth.go - fixme
package api

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	signKey    *rsa.PrivateKey
	serverPort int
)

func createToken(user string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"level1",
		CustomerInfo{user, "human"},
	}
	return t.SignedString(signKey)
}

// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}
