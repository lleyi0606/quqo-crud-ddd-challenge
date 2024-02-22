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

	repoAuthorization := authorization.NewAuthorizatiionRepository()
	key := []byte(os.Getenv("DATABASE_URL"))

	ts, tErr := repoAuthorization.GenerateToken(key, int64(cus.CustomerID), cus.Password)
	if tErr != nil {
		return "", cus, tErr
	}

	return ts, cus, nil
}

// func (u *authorizationApp) Logout(c *gin.Context) {
// 	//check is the user is authenticated first
// 	metadata, err := au.tk.ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}
// 	//if the access token exist and it is still valid, then delete both the access token and the refresh token
// 	deleteErr := au.rd.DeleteTokens(metadata)
// 	if deleteErr != nil {
// 		c.JSON(http.StatusUnauthorized, deleteErr.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, "Successfully logged out")
// }
