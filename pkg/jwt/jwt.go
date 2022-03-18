package jwt

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	allowIssuer = "theSHduckuser-service"
)

// edit your jwt claims
type JwtClaims struct {
	Role []string
	jwt.StandardClaims
}

func isTrustIssuer(claims *JwtClaims) bool {
	return claims.Issuer == allowIssuer
}

func myKeyFunc(token *jwt.Token) (interface{}, error) {
	// can change to switch case if you allow multiple sign algorithm
	if token.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
	}
	return []byte(os.Getenv("TOKEN_SECRET")), nil
}

func ReadTokenClaims(token *jwt.Token) (*JwtClaims, error) {
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("type of token's claims not match")
	}
	return claims, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, myKeyFunc)
	if err != nil {
		return nil, err
	}
	claims, err := ReadTokenClaims(token)
	if err != nil {
		return nil, err
	}
	if !isTrustIssuer(claims) {
		return nil, fmt.Errorf("token not from trust issuer")
	}
	return token, nil
}
