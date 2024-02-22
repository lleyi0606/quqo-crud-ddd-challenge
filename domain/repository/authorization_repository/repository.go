package authorization_repository

import (
	entity "products-crud/domain/entity/customer_entity"

	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationRepository interface {
	// Login(*entity.Customer) error
	// Logout() error
	// Refresh() error
	GenerateToken([]byte, int64, string) (string, error)

	ValidateToken(string, string) (*jwt.Token, error)
}

type AuthorizationHandlerRepository interface {
	Login(*entity.Customer) (string, *entity.Customer, error)
	// Logout() error
	// Refresh() error
}
