package webtoken

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mWO = "Hello user! You are an successfully authorized!"

//JWTCreate to create a new token
func JWTCreate(userID uint, accessLevel string, hours time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &jwt.StandardClaims{
		Id:        fmt.Sprint(userID),
		Subject:   accessLevel,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * hours).Unix(),
	})

	tokenstring, err := token.SignedString([]byte(mWO))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

//JWTParse Parse token
func JWTParse(tokenstring string) (jwt.StandardClaims, error) {
	tokenCustomClaims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenstring, &tokenCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mWO), nil
	})

	if err != nil {
		return tokenCustomClaims, err
	}

	return tokenCustomClaims, nil
}
