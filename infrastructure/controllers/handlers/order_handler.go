package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	repo        repository.OrderHandlerRepository
	Persistence *base.Persistence
}

// Orders constructor
func NewOrderController(p *base.Persistence) *OrderHandler {
	return &OrderHandler{
		Persistence: p,
	}
}

// @Summary Add order
// @Description Add an order to the database
// @Tags Order
// @Accept json
// @Produce json
// @Param order body entity.OrderInput true "Order data"
// @Success 201 {object} response_entity.Response "Order created"
// @Failure 400 {object} response_entity.Response "Invalid order_id format, Unable to parse form data, Unable to get Order from form"
// @Failure 500 {object} response_entity.Response "Application AddOrder error"
// @Router /orders [post]
func (p *OrderHandler) AddOrder(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// tracer := otel.Tracer("your-tracer")

	// // Start the main span for the HTTP request
	// _, span := tracer.Start(c, "http-request")
	// defer span.End()

	var order entity.OrderInput
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	cusIDString := c.GetString("userID")
	// Convert string to int64
	cusID, err := strconv.ParseUint(cusIDString, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails
		log.Println("Error converting cusIDString to int64:", err)
	} else {
		// Now cusID is of type uint64
		order.CustomerID = cusID
	}

	log.Print("id from context: ", cusID)
	// if exists {
	// 	if userID, ok := cusID.(uint64); ok {
	// 		log.Print("customer ID set from context", cusID)
	// 		order.CustomerID = userID
	// 	}
	// }

	p.repo = application.NewOrderApplication(p.Persistence)
	newOrder, err := p.repo.AddOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Order created.", newOrder))
}

// @Summary Get order
// @Description Get Order details by order_id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "order_id"
// @Success 200 {object} response_entity.Response "Successfully get Order"
// @Failure 400 {object} response_entity.Response "Invalid order_id GetOrder"
// @Failure 500 {object} response_entity.Response "Application GetOrder error"
// @Router /orders/{id} [get]
func (p *OrderHandler) GetOrder(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract order_id from the URL parameter
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid order_id GetOrder", ""))

		return
	}

	// Call the service to get a single Order by ID
	p.repo = application.NewOrderApplication(p.Persistence)
	order, err := p.repo.GetOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get orders.", order))

}

// @Summary Update order
// @Description Update an order in the database by ID
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "order_id"
// @Param order body entity.OrderInput true "Order data"
// @Success 201 {object} response_entity.Response "Order updated"
// @Failure 400 {object} response_entity.Response "Invalid order_id"
// @Failure 500 {object} response_entity.Response "Application UpdateOrder error"
// @Router /orders/{id} [put]
func (p *OrderHandler) UpdateOrder(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid order_id UpdateOrder", ""))
		return
	}

	p.repo = application.NewOrderApplication(p.Persistence)
	cus, _ := p.repo.GetOrder(orderID)

	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for Order update: %+v", cus)

	cus.OrderID = orderID

	newOrder, err := p.repo.UpdateOrder(cus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Order updated. ", newOrder))
}

// @Summary Delete order
// @Description Delete an Order from the database by ID
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "order_id"
// @Success 200 {object} response_entity.Response "Order deleted"
// @Failure 400 {object} response_entity.Response "Invalid order_id DeleteOrder"
// @Failure 500 {object} response_entity.Response "Application DeleteOrder error"
// @Router /orders/{id} [delete]
func (p *OrderHandler) DeleteOrder(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract order_id from the URL parameter
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid order_id DeleteOrder", ""))
		return
	}

	// Call the service to get a single Order by ID
	p.repo = application.NewOrderApplication(p.Persistence)
	err = p.repo.DeleteOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single Order
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Order deleted.", ""))
}
