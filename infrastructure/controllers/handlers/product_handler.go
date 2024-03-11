package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/product_entity"
	repository "products-crud/domain/repository/product_respository"
	loggerOpt "products-crud/infrastructure/implementations/logger"
	base "products-crud/infrastructure/persistences"

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

	// info := loggerentity.FunctionInfo{
	// 	FunctionName: "AddProduct",
	// 	Path:         "infrastructure/handlers/",
	// 	Description:  "Handles JSON input of AddProduct",
	// 	Body:         nil,
	// }

	// logger, endFunc := logger.NewLoggerRepositories(p.Persistence, c, info, []string{"Honeycomb", "zap"}, logger.SetNewOtelContext())
	// defer endFunc()

	logger := p.Persistence.Logger
	logger.Start(c, "infrastructure/handlers/AddProduct", map[string]interface{}{}, loggerOpt.SetNewOtelContext())
	defer p.Persistence.Logger.End()

	responseContextData := response_entity.ResponseContext{Ctx: c}

	var pdt entity.ProductWithStockAndWarehouse
	if err := c.ShouldBindJSON(&pdt); err != nil {
		logger.Error("invalid json", map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	pdt.ProductID = entity.GenerateProductID()
	log.Print("product id is ", pdt.ProductID)

	logger.Info("add product in handlers", map[string]interface{}{"input": pdt})

	p.p_repo = application.NewProductApplication(p.Persistence, c)
	newProduct, err := p.p_repo.AddProduct(&pdt)
	if err != nil {
		logger.Error(err.Error(), map[string]interface{}{"error": err})
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

	p.p_repo = application.NewProductApplication(p.Persistence, c)
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

// @Summary Retrieve all products for users
// @Description Retrieve all products from the database
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response_entity.Response "Product getted"
// @Failure 500 {object} response_entity.Response "Application GetProducts error"
// @Router /products [get]
func (p *ProductHandler) GetProductsUser(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var products []entity.ProductUser
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence, c)
	products, err = p.p_repo.GetProductsUser()
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
// @Param id path int true "product_id"
// @Success 200 {object} response_entity.Response "Product retrieved"
// @Failure 400 {object} response_entity.Response "Invalid product_id"
// @Failure 500 {object} response_entity.Response "Application GetProduct error"
// @Router /products/{id} [get]
func (p *ProductHandler) GetProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product_id from the URL parameter
	productID := c.Param("id")
	// productID, err := strconv.ParseUint(productIDStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product_id GetProduct", ""))
	// 	return
	// }

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence, c)
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
// @Param id path int true "product_id"
// @Success 201 {object} response_entity.Response "Product updated"
// @Failure 400 {object} response_entity.Response "Invalid product_id"
// @Failure 500 {object} response_entity.Response "Application UpdateProduct error"
// @Router /products/{id} [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	productID := c.Param("id")

	p.p_repo = application.NewProductApplication(p.Persistence, c)
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
// @Param id path int true "product_id"
// @Success 200 {object} response_entity.Response "Product deleted"
// @Failure 400 {object} response_entity.Response "Invalid product_id"
// @Failure 500 {object} response_entity.Response "Application DeleteProduct error"
// @Router /products/{id} [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product_id from the URL parameter
	productID := c.Param("id")
	// productID, err := strconv.ParseUint(productIDStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product_id DeleteProduct", ""))
	// 	return
	// }

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence, c)
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

	// info := loggerentity.FunctionInfo{
	// 	FunctionName: "SearchProducts",
	// 	Path:         "/infrastructure/handlers/",
	// 	Description:  "Gets query keyword for search",
	// 	Body:         nil,
	// }
	// logger, endFunc := logger.NewLoggerRepositories(p.Persistence, c, info, []string{"Honeycomb", "zap"})
	// defer endFunc()

	keyword := c.Query("name")

	var products []entity.Product
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence, c)
	products, err = p.p_repo.SearchProducts(keyword)
	if err != nil {
		// logger.Error(err.Error(), map[string]interface{}{"query": keyword})
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

	p.p_repo = application.NewProductApplication(p.Persistence, c)
	for _, pdt := range pdts {
		_, err := p.p_repo.AddProduct(&pdt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
			return
		}
	}

	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Products added", ""))
}
