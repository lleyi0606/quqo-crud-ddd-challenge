package authorization

import (
	"errors"
	"fmt"
	"log"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/authorization_repository"
	"products-crud/infrastructure/implementations/cache"
	base "products-crud/infrastructure/persistences"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authorizationRepo struct {
	p *base.Persistence
}

func NewAuthorizationRepository(p *base.Persistence) repository.AuthorizationRepository {
	return &authorizationRepo{p}
}

func (a authorizationRepo) GenerateToken(key []byte, userId int64, credential string) (string, error) {

	//new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["user_id"] = fmt.Sprint(userId)
	claims["credential"] = credential
	claims["exp"] = time.Now().Add(time.Hour*240).UnixNano() / int64(time.Millisecond)

	log.Println("in generate token: ", userId, claims["user_id"])
	//Set user roles
	//claims["roles"] = roles

	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func (a authorizationRepo) ValidateToken(tokenString string, key string) (*jwt.Token, error) {

	isBlacklisted, _ := a.IsTokenInBlacklist(tokenString)
	if isBlacklisted {
		return nil, errors.New("token is blacklisted")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}

func (a authorizationRepo) AddTokenToBlacklist(tokenString string) error {

	blacklistedToken := redis_entity.BlacklistedToken{
		Token:       tokenString,
		Reason:      "User logout", // Optional: Reason for revocation
		RevokedByID: "TEMP_DUMMY",  // Optional: ID of the user who revoked the token
	}

	cacheRepo := cache.NewCacheRepository(a.p, "redis")
	err := cacheRepo.SetKey(fmt.Sprintf("%s%s", redis_entity.RedisJWTData, strings.TrimPrefix(tokenString, "Bearer ")), blacklistedToken, redis_entity.RedisExpirationJwt)

	return err
}

func (a authorizationRepo) IsTokenInBlacklist(tokenString string) (bool, error) {

	var token *redis_entity.BlacklistedToken
	cacheRepo := cache.NewCacheRepository(a.p, "redis")
	err := cacheRepo.GetKey(fmt.Sprintf("%s%s", redis_entity.RedisJWTData, tokenString), &token)

	if token == nil {
		return false, err
	}

	return true, err
}
