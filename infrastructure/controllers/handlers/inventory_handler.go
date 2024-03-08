package handlers

import (
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/inventory_entity"
	repository "products-crud/domain/repository/inventory_respository"
	base "products-crud/infrastructure/persistences"

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

// @Summary Get inventory
// @Description Get inventory details by product ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.Inventory "Inventory details"
// @Failure 400 {object} response_entity.Response "Invalid product ID GetInventory"
// @Failure 500 {object} response_entity.Response "Application GetInventory error"
// @Router /inventories/{id} [get]
func (p *InventoryHandler) GetInventory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	// productID, err := strconv.ParseUint(productIDStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID GetInventory", ""))

	// 	return
	// }

	// Call the service to get a single product by ID
	p.p_repo = application.NewInventoryApplication(p.Persistence)
	product, err := p.p_repo.GetInventory(productIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get inventory. ", product))
}

// @Summary Update stock
// @Description Update stock details for a product by ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param Object body entity.InventoryStockOnly true "Stock details to update"
// @Success 201 {object} response_entity.Response "Stock updated"
// @Failure 400 {object} response_entity.Response "Invalid product ID UpdateStock or Invalid JSON"
// @Failure 500 {object} response_entity.Response "Application UpdateStock error"
// @Router /inventories/{id} [put]
func (p *InventoryHandler) UpdateStock(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	productIDStr := c.Param("id")
	// productID, err := strconv.ParseUint(productIDStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID UpdateStock", ""))
	// 	return
	// }

	var ivt entity.InventoryStockOnly

	p.p_repo = application.NewInventoryApplication(p.Persistence)

	if err := c.ShouldBindJSON(&ivt); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid json", ""))
		return
	}

	newProduct, err := p.p_repo.UpdateStock(productIDStr, &ivt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Stock updated.", newProduct))
}
