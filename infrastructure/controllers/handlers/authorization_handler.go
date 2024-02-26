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

// @Summary User login
// @Description Log in with user credentials and obtain an access token.
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body body entity.Customer true "User credentials for login"
// @Success 200 {object} response_entity.Response "Successful login"
// @Failure 422 {object} response_entity.Response "Invalid JSON provided"
// @Failure 500 {object} response_entity.Response "Internal Server Error"
// @Router /login [post]
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
	userData["id"] = user.CustomerID
	userData["name"] = user.Name

	c.Header("Authorization", "Bearer "+ts)

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Log in success.", userData))
}

// @Summary User logout
// @Description Log out and invalidate the user token.
// @Tags Authorization
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <token>) "Access token for authentication"
// @Success 200 {object} response_entity.Response "Successful logout"
// @Failure 500 {object} response_entity.Response "Internal Server Error"
// @Router /logout [post]
func (ah *AuthorizationHandler) Logout(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	authorizationHeader := c.Request.Header.Get("Authorization")
	ah.repo = application.NewAuthorizationApplication(ah.Persistence)
	err := ah.repo.Logout(authorizationHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Log out success.", ""))

}
