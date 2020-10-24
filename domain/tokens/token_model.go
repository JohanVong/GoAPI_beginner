package tokens

import (
	"time"

	"github.com/JohanVong/GoAPI_beginner/utils/errors"
	"github.com/JohanVong/GoAPI_beginner/utils/webtoken"
	jwt "github.com/dgrijalva/jwt-go"
)

// Token model
type Token struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

//ValidateToken func validates a token
func (token *Token) ValidateToken() *errors.CustomError {
	var err error
	var tokenClaims jwt.StandardClaims

	tokenClaims, err = webtoken.JWTParse(token.Token)
	if err != nil {
		return errors.TextError("An error on JSON web token parsing occurred", err.Error())
	}

	if time.Now().Unix() >= tokenClaims.ExpiresAt {
		return errors.TextError("Token is obsolete", "Invalid token error")
	}

	return nil
}
