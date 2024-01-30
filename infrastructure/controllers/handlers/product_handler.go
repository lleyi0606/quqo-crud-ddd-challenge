package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
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
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Get products.", products))
}

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

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID UpdateProduct", ""))
		return
	}

	// var pdt entity.Product

	p.p_repo = application.NewProductApplication(p.Persistence)
	pdt, _ := p.p_repo.GetProduct(productID)

	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// var newPdt entity.ProductUpdate
	// newPdt = entity.SqlProductRToProductForUpdate(pdt, productID)

	// Log the JSON input
	log.Printf("Received JSON input for product update: %+v", pdt)

	pdt.ProductID = productID

	// p.p_repo = application.NewProductApplication(p.Persistence)
	newProduct, err := p.p_repo.UpdateProduct(pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Product updated. ", newProduct))
}

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
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Product searched.", products))
}

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
