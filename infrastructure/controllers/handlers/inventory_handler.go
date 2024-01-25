package handlers

import (
	"net/http"
	"products-crud/application"
	entity "products-crud/domain/entity/inventory_entity"
	repository "products-crud/domain/repository/inventory_respository"
	base "products-crud/infrastructure/persistences"
	"strconv"

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

// func (p *InventoryHandler) AddInventory(c *gin.Context) {
// 	var ivt entity.Inventory
// 	if err := c.ShouldBindJSON(&ivt); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"invalid_json": "invalid json",
// 		})
// 		return
// 	}

// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	newProduct, err := p.p_repo.AddInventory(&ivt)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	c.JSON(http.StatusCreated, newProduct)
// }

// func (p *InventoryHandler) GetInventories(c *gin.Context) {
// 	var products []entity.Inventory
// 	var err error

// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	products, err = p.p_repo.GetInventories()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, products)
// }

func (p *InventoryHandler) GetInventory(c *gin.Context) {
	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID GetInventory"})
		return
	}

	// Call the service to get a single product by ID
	p.p_repo = application.NewInventoryApplication(p.Persistence)
	product, err := p.p_repo.GetInventory(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, product)
}

func (p *InventoryHandler) UpdateStock(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID UpdateStock"})
		return
	}

	var ivt entity.InventoryStockOnly

	p.p_repo = application.NewInventoryApplication(p.Persistence)
	// ivtFull, _ := p.p_repo.GetInventory(productID)

	if err := c.ShouldBindJSON(&ivt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	// p.p_repo = application.NewInventoryApplication(p.Persistence)
	newProduct, err := p.p_repo.UpdateStock(productID, &ivt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// func (p *InventoryHandler) UpdateInventory(c *gin.Context) {
// 	productIDStr := c.Param("id")
// 	productID, err := strconv.ParseUint(productIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID UpdateProduct"})
// 		return
// 	}

// 	// var ivt entity.Inventory

// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	ivt, _ := p.p_repo.GetInventory(productID)

// 	if err := c.ShouldBindJSON(&ivt); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"invalid_json": "invalid json",
// 		})
// 		return
// 	}

// 	// var newPdt entity.ProductUpdate
// 	// newPdt = entity.SqlProductRToProductForUpdate(ivt, productID)

// 	// Log the JSON input
// 	log.Printf("Received JSON input for product update: %+v", ivt)

// 	ivt.InventoryID = productID

// 	// p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	newProduct, err := p.p_repo.UpdateInventory(ivt)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	c.JSON(http.StatusCreated, newProduct)
// }

// func (p *InventoryHandler) DeleteInventory(c *gin.Context) {
// 	// Extract product ID from the URL parameter
// 	productIDStr := c.Param("id")
// 	productID, err := strconv.ParseUint(productIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID DeleteProduct"})
// 		return
// 	}

// 	// Call the service to get a single product by ID
// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	product, err := p.p_repo.DeleteInventory(productID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	// Respond with the single product
// 	c.JSON(http.StatusOK, product)
// }

// func (p *InventoryHandler) SearchInventory(c *gin.Context) {
// 	keyword := c.Query("name")

// 	var products []entity.Inventory
// 	var err error

// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	products, err = p.p_repo.SearchInventory(keyword)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, products)
// }

// func (p *InventoryHandler) AddInventories(c *gin.Context) {
// 	var pdts []entity.Inventory
// 	if err := c.ShouldBindJSON(&pdts); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"invalid_json": "invalid json",
// 		})
// 		return
// 	}

// 	p.p_repo = application.NewInventoryApplication(p.Persistence)
// 	for _, ivt := range pdts {
// 		_, err := p.p_repo.AddInventory(&ivt)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, err)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusCreated, pdts)
// }
