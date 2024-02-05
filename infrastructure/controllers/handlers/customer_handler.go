package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/customer_entity"
	repository "products-crud/domain/repository/customer_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	repo        repository.CustomerHandlerRepository
	Persistence *base.Persistence
}

// Customers constructor
func NewCustomerController(p *base.Persistence) *CustomerHandler {
	return &CustomerHandler{
		Persistence: p,
	}
}

// @Summary Add Customer
// @Description Add an Customer to the database
// @Tags Customer
// @Accept mpfd
// @Produce json
// @Param Customer_id formData int64 true "Customer ID"
// @Param caption formData string false "Caption"
// @Param Customer_file formData file true "Customer file"
// @Success 201 {object} response_entity.Response "Customer created"
// @Failure 400 {object} response_entity.Response "Invalid Customer ID format, Unable to parse form data, Unable to get Customer from form"
// @Failure 500 {object} response_entity.Response "Application AddCustomer error"
// @Router /Customers [post]
func (p *CustomerHandler) AddCustomer(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var cus entity.Customer
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	p.repo = application.NewCustomerApplication(p.Persistence)
	newCustomer, err := p.repo.AddCustomer(&cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Customer created.", newCustomer))
}

// @Summary Get customers
// @Description Get Customer details by Customer ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} response_entity.Response "Successfully get Customers"
// @Failure 400 {object} response_entity.Response "Invalid Customer ID GetCustomer"
// @Failure 500 {object} response_entity.Response "Application GetCustomer error"
// @Router /Customers/{id} [get]
func (p *CustomerHandler) GetCustomer(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract Customer ID from the URL parameter
	customerIDStr := c.Param("id")
	customerID, err := strconv.ParseUint(customerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid Customer ID GetCustomer", ""))

		return
	}

	// Call the service to get a single Customer by ID
	p.repo = application.NewCustomerApplication(p.Persistence)
	customer, err := p.repo.GetCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get customers.", customer))

}

// @Summary Update customer
// @Description Update a Customer in the database by ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 201 {object} response_entity.Response "Customer updated"
// @Failure 400 {object} response_entity.Response "Invalid Customer ID"
// @Failure 500 {object} response_entity.Response "Application UpdateCustomer error"
// @Router /Customers/{id} [put]
func (p *CustomerHandler) UpdateCustomer(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	customerIDStr := c.Param("id")
	customerID, err := strconv.ParseUint(customerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid Customer ID UpdateCustomer", ""))
		return
	}

	p.repo = application.NewCustomerApplication(p.Persistence)
	cus, _ := p.repo.GetCustomer(customerID)

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for Customer update: %+v", cus)

	cus.CustomerID = customerID

	newCustomer, err := p.repo.UpdateCustomer(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Customer updated. ", newCustomer))
}

// @Summary Delete Customer
// @Description Delete an Customer from the database by ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} response_entity.Response "Customer deleted"
// @Failure 400 {object} response_entity.Response "Invalid Customer ID DeleteCustomer"
// @Failure 500 {object} response_entity.Response "Application DeleteCustomer error"
// @Router /Customers/{id} [delete]
func (p *CustomerHandler) DeleteCustomer(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract Customer ID from the URL parameter
	customerIDStr := c.Param("id")
	customerID, err := strconv.ParseUint(customerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid Customer ID DeleteCustomer", ""))
		return
	}

	// Call the service to get a single Customer by ID
	p.repo = application.NewCustomerApplication(p.Persistence)
	err = p.repo.DeleteCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single Customer
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Customer deleted.", ""))
}
