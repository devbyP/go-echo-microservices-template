package jwt

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go/v4"
)

// edit your jwt claims
type MyJwtClaims struct {
	jwt.StandardClaims
}

const (
	allowIssuer = "theSHduckuser-service"
)

func MyKeyFunc(token *jwt.Token) (interface{}, error) {
	// can change to switch case if you allow multiple sign algorithm
	if token.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("error in claims type assertion")
	}
	if claims.Issuer != allowIssuer {
		return nil, fmt.Errorf("not from trust source")
	}
	if claims.Subject == "" {
		return nil, fmt.Errorf("no user id")
	}
	return []byte(os.Getenv("TOKEN_SECRET")), nil
}
