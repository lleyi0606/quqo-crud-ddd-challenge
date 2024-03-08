package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/warehouse_entity"
	repository "products-crud/domain/repository/warehouse_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	repo        repository.WarehouseHandlerRepository
	Persistence *base.Persistence
}

// warehouses constructor
func NewWarehouseController(p *base.Persistence) *WarehouseHandler {
	return &WarehouseHandler{
		Persistence: p,
	}
}

// @Summary Add warehouse
// @Description Add an Warehouse to the database
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param warehouse body entity.Warehouse true "Warehouse data"
// @Success 201 {object} response_entity.Response "Warehouse created"
// @Failure 400 {object} response_entity.Response "Invalid warehouse_id format, Unable to parse form data, Unable to get Warehouse from form"
// @Failure 500 {object} response_entity.Response "Application AddWarehouse error"
// @Router /warehouses [post]
func (p *WarehouseHandler) AddWarehouse(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var cus entity.Warehouse
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	p.repo = application.NewWarehouseApplication(p.Persistence)
	newWarehouse, err := p.repo.AddWarehouse(&cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Warehouse created.", newWarehouse))
}

// @Summary Get warehouses
// @Description Get Warehouse details by warehouse_id
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "warehouse_id"
// @Success 200 {object} response_entity.Response "Successfully get warehouses"
// @Failure 400 {object} response_entity.Response "Invalid warehouse_id GetWarehouse"
// @Failure 500 {object} response_entity.Response "Application GetWarehouse error"
// @Router /warehouses/{id} [get]
func (p *WarehouseHandler) GetWarehouse(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract warehouse_id from the URL parameter
	warehouseIDStr := c.Param("id")
	warehouseID, err := strconv.ParseUint(warehouseIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid warehouse_id GetWarehouse", ""))

		return
	}

	// Call the service to get a single Warehouse by ID
	p.repo = application.NewWarehouseApplication(p.Persistence)
	warehouse, err := p.repo.GetWarehouse(warehouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get warehouses.", warehouse))
}

// @Summary Get all warehouses
// @Description Get all Warehouse details
// @Tags Warehouse
// @Accept json
// @Produce json
// @Success 200 {object} response_entity.Response "Successfully get warehouses"
// @Failure 500 {object} response_entity.Response "Application GetAllWarehouse error"
// @Router /warehouses/{id} [get]
func (p *WarehouseHandler) GetWarehouses(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Call the service to get a single Warehouse by ID
	p.repo = application.NewWarehouseApplication(p.Persistence)
	warehouse, err := p.repo.GetWarehouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get warehouses.", warehouse))
}

// @Summary Update warehouse
// @Description Update a Warehouse in the database by ID
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "warehouse_id"
// @Param warehouse body entity.Warehouse true "Warehouse data"
// @Success 201 {object} response_entity.Response "Warehouse updated"
// @Failure 400 {object} response_entity.Response "Invalid warehouse_id"
// @Failure 500 {object} response_entity.Response "Application UpdateWarehouse error"
// @Router /warehouses/{id} [put]
func (p *WarehouseHandler) UpdateWarehouse(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	warehouseIDStr := c.Param("id")
	warehouseID, err := strconv.ParseUint(warehouseIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid warehouse_id UpdateWarehouse", ""))
		return
	}

	p.repo = application.NewWarehouseApplication(p.Persistence)
	cus, _ := p.repo.GetWarehouse(warehouseID)

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for Warehouse update: %+v", cus)

	cus.WarehouseID = warehouseID

	newWarehouse, err := p.repo.UpdateWarehouse(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Warehouse updated. ", newWarehouse))
}

// @Summary Delete warehouse
// @Description Delete a warehouse from the database by ID
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "warehouse_id"
// @Success 200 {object} response_entity.Response "Warehouse deleted"
// @Failure 400 {object} response_entity.Response "Invalid warehouse_id DeleteWarehouse"
// @Failure 500 {object} response_entity.Response "Application DeleteWarehouse error"
// @Router /warehouses/{id} [delete]
func (p *WarehouseHandler) DeleteWarehouse(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract warehouse_id from the URL parameter
	warehouseIDStr := c.Param("id")
	warehouseID, err := strconv.ParseUint(warehouseIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid warehouse_id DeleteWarehouse", ""))
		return
	}

	// Call the service to get a single Warehouse by ID
	p.repo = application.NewWarehouseApplication(p.Persistence)
	err = p.repo.DeleteWarehouse(warehouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single Warehouse
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Warehouse deleted.", ""))
}
