package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/product_entity"
	repository "products-crud/domain/repository/product_respository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	p_repo      repository.ProductHandlerRepository
	Persistence *base.Persistence
}

// Products constructor
func NewProductController(p *base.Persistence) *ProductHandler {
	return &ProductHandler{
		Persistence: p,
	}
}

// @Summary Create new product
// @Description Adds a new product to the database
// @Tags Product
// @Accept json
// @Produce json
// @Param body body entity.ProductWithStockAndWarehouse true "Entity of the body request"
// @Success 201 {object} response_entity.Response "Product created"
// @Failure 422 {object} response_entity.Response "Parse input error"
// @Failure 500 {object} response_entity.Response "Application AddProduct error"
// @Router /products [post]
func (p *ProductHandler) AddProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var pdt entity.ProductWithStockAndWarehouse
	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	p.p_repo = application.NewProductApplication(p.Persistence)
	newProduct, err := p.p_repo.AddProduct(&pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Product created.", newProduct))
}

// @Summary Retrieve all products
// @Description Retrieve all products from the database
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response_entity.Response "Product getted"
// @Failure 500 {object} response_entity.Response "Application GetProducts error"
// @Router /products [get]
func (p *ProductHandler) GetProducts(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var products []entity.Product
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence)
	products, err = p.p_repo.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get products.",
		map[string]interface{}{
			"result": products,
		},
	))
}

// @Summary Retrieve a product
// @Description Retrieve a product from the database by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response_entity.Response "Product retrieved"
// @Failure 400 {object} response_entity.Response "Invalid product ID"
// @Failure 500 {object} response_entity.Response "Application GetProduct error"
// @Router /products/{id} [get]
func (p *ProductHandler) GetProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID GetProduct", ""))
		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence)
	product, err := p.p_repo.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get product.", product))
}

// @Summary Update a product
// @Description Update a product in the database by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 201 {object} response_entity.Response "Product updated"
// @Failure 400 {object} response_entity.Response "Invalid product ID"
// @Failure 500 {object} response_entity.Response "Application UpdateProduct error"
// @Router /products/{id} [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID UpdateProduct", ""))
		return
	}

	p.p_repo = application.NewProductApplication(p.Persistence)
	pdt, _ := p.p_repo.GetProduct(productID)

	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for product update: %+v", pdt)

	pdt.ProductID = productID

	newProduct, err := p.p_repo.UpdateProduct(pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Product updated. ", newProduct))
}

// @Summary Delete a product
// @Description Delete a product from the database by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response_entity.Response "Product deleted"
// @Failure 400 {object} response_entity.Response "Invalid product ID"
// @Failure 500 {object} response_entity.Response "Application DeleteProduct error"
// @Router /products/{id} [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID DeleteProduct", ""))
		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence)
	product, err := p.p_repo.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Product deleted", product))
}

// @Summary Search for products
// @Description Search for products by keyword
// @Tags Product
// @Accept json
// @Produce json
// @Param name query string true "Search keyword for product name"
// @Success 200 {object} response_entity.Response "Products searched"
// @Failure 500 {object} response_entity.Response "Application SearchProducts error"
// @Router /products/search [get]
func (p *ProductHandler) SearchProducts(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	keyword := c.Query("name")

	var products []entity.Product
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence)
	products, err = p.p_repo.SearchProducts(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Product searched.",
		map[string]interface{}{
			"result": products,
		}))
}

// @Summary Add products
// @Description Add multiple products to the database
// @Tags Product
// @Accept json
// @Produce json
// @Param Object body []entity.ProductWithStockAndWarehouse true "Array of products to add"
// @Success 201 {object} response_entity.Response "Products added"
// @Failure 422 {object} response_entity.Response "Parse input error"
// @Router /products [post]
func (p *ProductHandler) AddProducts(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var pdts []entity.ProductWithStockAndWarehouse
	if err := c.ShouldBindJSON(&pdts); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	p.p_repo = application.NewProductApplication(p.Persistence)
	for _, pdt := range pdts {
		_, err := p.p_repo.AddProduct(&pdt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
			return
		}
	}

	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Products added", ""))
}
