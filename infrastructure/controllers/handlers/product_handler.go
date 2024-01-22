package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
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
	var pdt entity.Product
	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	p.p_repo = application.NewProductApplication(p.Persistence)
	newProduct, err := p.p_repo.AddProduct(&pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

func (p *ProductHandler) GetProducts(c *gin.Context) {
	var products []entity.Product
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence)
	products, err = p.p_repo.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductHandler) GetProduct(c *gin.Context) {
	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID GetProduct"})
		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence)
	product, err := p.p_repo.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID UpdateProduct"})
		return
	}

	// var pdt entity.Product

	p.p_repo = application.NewProductApplication(p.Persistence)
	pdt, _ := p.p_repo.GetProduct(productID)

	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	// var newPdt entity.ProductUpdate
	// newPdt = entity.SqlProductRToProductForUpdate(pdt, productID)

	// Log the JSON input
	log.Printf("Received JSON input for product update: %+v", pdt)

	pdt.ID = productID

	// p.p_repo = application.NewProductApplication(p.Persistence)
	newProduct, err := p.p_repo.UpdateProduct(pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID DeleteProduct"})
		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewProductApplication(p.Persistence)
	product, err := p.p_repo.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) SearchProducts(c *gin.Context) {
	keyword := c.Query("name")

	var products []entity.Product
	var err error

	p.p_repo = application.NewProductApplication(p.Persistence)
	products, err = p.p_repo.SearchProducts(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductHandler) AddProducts(c *gin.Context) {
	var pdts []entity.Product
	if err := c.ShouldBindJSON(&pdts); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	p.p_repo = application.NewProductApplication(p.Persistence)
	for _, pdt := range pdts {
		_, err := p.p_repo.AddProduct(&pdt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusCreated, pdts)
}
