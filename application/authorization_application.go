package application

import (
	"os"
	entity "products-crud/domain/entity/customer_entity"
	repository "products-crud/domain/repository/authorization_repository"
	"products-crud/infrastructure/implementations/authorization"
	"products-crud/infrastructure/implementations/customer"
	base "products-crud/infrastructure/persistences"
)

type authorizationApp struct {
	p *base.Persistence
}

func NewAuthorizationApplication(p *base.Persistence) repository.AuthorizationHandlerRepository {
	return &authorizationApp{p}
}

func (u *authorizationApp) Login(user *entity.Customer) (string, *entity.Customer, error) {
	// return repoAuthorization.AddAuthorization(cat)

	repoCustomer := customer.NewCustomerRepository(u.p)
	cus, userErr := repoCustomer.GetCustomerByUsernameAndPassword(user)
	if userErr != nil {
		return "", nil, userErr
	}

	repoAuthorization := authorization.NewAuthorizationRepository(u.p)

	key := []byte(os.Getenv("DATABASE_URL"))

	ts, tErr := repoAuthorization.GenerateToken(key, int64(cus.CustomerID), cus.Password)
	if tErr != nil {
		return "", cus, tErr
	}

	return ts, cus, nil
}

func (u *authorizationApp) Logout(tokenString string) error {
	repoAuthorization := authorization.NewAuthorizationRepository(u.p)

	err := repoAuthorization.AddTokenToBlacklist(tokenString)
	return err
}
