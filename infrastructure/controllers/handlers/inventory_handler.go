package handlers

import (
	"net/http"
	"products-crud/application"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/inventory_entity"
	repository "products-crud/domain/repository/inventory_respository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	p_repo      repository.InventoryHandlerRepository
	Persistence *base.Persistence
}

// Products constructor
func NewInventoryController(p *base.Persistence) *InventoryHandler {
	return &InventoryHandler{
		Persistence: p,
	}
}

func (p *InventoryHandler) GetInventory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID GetInventory", ""))

		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewInventoryApplication(p.Persistence)
	product, err := p.p_repo.GetInventory(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get inventory. ", product))
}

func (p *InventoryHandler) UpdateStock(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID UpdateStock", ""))
		return
	}

	var ivt entity.InventoryStockOnly

	p.p_repo = application.NewInventoryApplication(p.Persistence)
	// ivtFull, _ := p.p_repo.GetInventory(productID)

	if err := c.ShouldBindJSON(&ivt); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid json", ""))
		return
	}

	// p.p_repo = application.NewInventoryApplication(p.Persistence)
	newProduct, err := p.p_repo.UpdateStock(productID, &ivt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Stock updated.", newProduct))
}
