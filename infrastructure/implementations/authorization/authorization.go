package authorization

import (
	repository "products-crud/domain/repository/authorization_repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authorizationRepo struct {
}

func NewAuthorizatiionRepository() repository.AuthorizationRepository {
	return &authorizationRepo{}
}

func (a authorizationRepo) GenerateToken(key []byte, userId int64, credential string) (string, error) {

	//new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["user_id"] = userId
	claims["credential"] = credential
	claims["exp"] = time.Now().Add(time.Hour*720).UnixNano() / int64(time.Millisecond)

	//Set user roles
	//claims["roles"] = roles

	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func (a authorizationRepo) ValidateToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}
