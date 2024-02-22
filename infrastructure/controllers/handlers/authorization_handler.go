package handlers

import (
	"net/http"
	"products-crud/application"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/customer_entity"
	repository "products-crud/domain/repository/authorization_repository"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

type AuthorizationHandler struct {
	repo        repository.AuthorizationHandlerRepository
	Persistence *base.Persistence
}

// authorization constructor
func NewAuthorizationController(p *base.Persistence) *AuthorizationHandler {
	return &AuthorizationHandler{
		Persistence: p,
	}
}

func (ah *AuthorizationHandler) Login(c *gin.Context) {
	var cus *entity.Customer
	responseContextData := response_entity.ResponseContext{Ctx: c}

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	ah.repo = application.NewAuthorizationApplication(ah.Persistence)
	ts, user, err := ah.repo.Login(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	userData := make(map[string]interface{})
	userData["access_token"] = ts
	userData["id"] = user.ID
	userData["name"] = user.Name

	c.Header("Authorization", "Bearer "+ts)

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Log in success.", userData))
}
