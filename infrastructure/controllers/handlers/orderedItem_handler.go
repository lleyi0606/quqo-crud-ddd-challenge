package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/orderedItem_entity"
	repository "products-crud/domain/repository/orderedItem_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderedItemHandler struct {
	repo        repository.OrderedItemHandlerRepository
	Persistence *base.Persistence
}

// OrderedItems constructor
func NewOrderedItemController(p *base.Persistence) *OrderedItemHandler {
	return &OrderedItemHandler{
		Persistence: p,
	}
}

// @Summary Retrieve all ordered items
// @Description Retrieve all ordered items from the database
// @Tags Ordered Item
// @Accept json
// @Produce json
// @Success 200 {object} response_entity.Response "Ordered items getted"
// @Failure 500 {object} response_entity.Response "Application GetOrderedItems error"
// @Router /orderedItem [get]
func (p *OrderedItemHandler) GetOrderedItems(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var items []entity.OrderedItem
	var err error

	p.repo = application.NewOrderedItemApplication(p.Persistence)
	items, err = p.repo.GetOrderedItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get ordered items.",
		map[string]interface{}{
			"result": items,
		},
	))
}

// @Summary Get orderedItems
// @Description Get OrderedItem details by OrderedItem ID
// @Tags OrderedItem
// @Accept json
// @Produce json
// @Param id path int true "OrderedItem ID"
// @Success 200 {object} response_entity.Response "Successfully get OrderedItems"
// @Failure 400 {object} response_entity.Response "Invalid OrderedItem ID GetOrderedItem"
// @Failure 500 {object} response_entity.Response "Application GetOrderedItem error"
// @Router /OrderedItems/{id} [get]
func (p *OrderedItemHandler) GetOrderedItem(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract OrderedItem ID from the URL parameter
	orderedItemIDStr := c.Param("id")
	orderedItemID, err := strconv.ParseUint(orderedItemIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid OrderedItem ID GetOrderedItem", ""))

		return
	}

	// Call the service to get a single OrderedItem by ID
	p.repo = application.NewOrderedItemApplication(p.Persistence)
	orderedItem, err := p.repo.GetOrderedItem(orderedItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get orderedItems.", orderedItem))

}

// @Summary Update orderedItem
// @Description Update a OrderedItem in the database by ID
// @Tags OrderedItem
// @Accept json
// @Produce json
// @Param id path int true "OrderedItem ID"
// @Success 201 {object} response_entity.Response "OrderedItem updated"
// @Failure 400 {object} response_entity.Response "Invalid OrderedItem ID"
// @Failure 500 {object} response_entity.Response "Application UpdateOrderedItem error"
// @Router /OrderedItems/{id} [put]
func (p *OrderedItemHandler) UpdateOrderedItem(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	orderedItemIDStr := c.Param("id")
	orderedItemID, err := strconv.ParseUint(orderedItemIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid OrderedItem ID UpdateOrderedItem", ""))
		return
	}

	p.repo = application.NewOrderedItemApplication(p.Persistence)
	cus, _ := p.repo.GetOrderedItem(orderedItemID)

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for OrderedItem update: %+v", cus)

	// cus.OrderedItemID = orderedItemID

	newOrderedItem, err := p.repo.UpdateOrderedItem(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "OrderedItem updated. ", newOrderedItem))
}

// @Summary Delete OrderedItem
// @Description Delete an OrderedItem from the database by ID
// @Tags OrderedItem
// @Accept json
// @Produce json
// @Param id path int true "OrderedItem ID"
// @Success 200 {object} response_entity.Response "OrderedItem deleted"
// @Failure 400 {object} response_entity.Response "Invalid OrderedItem ID DeleteOrderedItem"
// @Failure 500 {object} response_entity.Response "Application DeleteOrderedItem error"
// @Router /OrderedItems/{id} [delete]
func (p *OrderedItemHandler) DeleteOrderedItem(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract OrderedItem ID from the URL parameter
	orderedItemIDStr := c.Param("id")
	orderedItemID, err := strconv.ParseUint(orderedItemIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid OrderedItem ID DeleteOrderedItem", ""))
		return
	}

	// Call the service to get a single OrderedItem by ID
	p.repo = application.NewOrderedItemApplication(p.Persistence)
	err = p.repo.DeleteOrderedItem(orderedItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single OrderedItem
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "OrderedItem deleted.", ""))
}
