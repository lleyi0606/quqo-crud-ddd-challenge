package handlers

import (
	"net/http"
	"products-crud/application"
	"products-crud/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	us application.ProductAppInterface
}

// Products constructor
func NewProducts(us application.ProductAppInterface) *Products {
	return &Products{
		us: us,
	}
}

func (p *Products) AddProduct(c *gin.Context) {
	var pdt entity.Product
	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	newProduct, err := p.us.AddProduct(&pdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

func (p *Products) GetProducts(c *gin.Context) {
	var products []entity.Product
	var err error
	products, err = p.us.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *Products) GetProduct(c *gin.Context) {
	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID GetProduct"})
		return
	}

	// Call the service to get a single product by ID
	product, err := p.us.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, product)
}

func (p *Products) UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID UpdateProduct"})
		return
	}

	var pdt entity.ProductToReceive
	if err := c.ShouldBindJSON(&pdt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	var newPdt entity.ProductUpdate
	newPdt = entity.SqlProductRToProductForUpdate(pdt, productID)

	newProduct, err := p.us.UpdateProduct(&newPdt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

func (p *Products) DeleteProduct(c *gin.Context) {
	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID DeleteProduct"})
		return
	}

	// Call the service to get a single product by ID
	product, err := p.us.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, product)
}

func (p *Products) SearchProducts(c *gin.Context) {
	keyword := c.Query("name")

	var products []entity.Product
	var err error
	products, err = p.us.SearchProducts(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}
