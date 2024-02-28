package handlers

import (
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
// @Router /ordereditems [get]
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
// @Description Get OrderedItem details by Order ID
// @Tags Ordered Item
// @Accept json
// @Produce json
// @Param id path int true "OrderedItem ID"
// @Success 200 {object} response_entity.Response "Successfully get OrderedItemsByOrderId"
// @Failure 400 {object} response_entity.Response "Invalid Order ID GetOrderedItemsByOrderId"
// @Failure 500 {object} response_entity.Response "Application GetOrderedItemsByOrderId error"
// @Router /ordereditems/{id} [get]
func (p *OrderedItemHandler) GetOrderedItemsByOrderId(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract OrderedItem ID from the URL parameter
	orderedItemIDStr := c.Param("id")
	orderedItemID, err := strconv.ParseUint(orderedItemIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid OrderedItem ID GetOrderedItem", ""))

		return
	}

	// Call the service to get OrderedItems by Order ID
	p.repo = application.NewOrderedItemApplication(p.Persistence)
	orderedItem, err := p.repo.GetOrderedItemsByOrderId(orderedItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get orderedItems.", orderedItem))

}
