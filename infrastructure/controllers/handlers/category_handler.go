package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	_ "products-crud/docs"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/category_entity"
	repository "products-crud/domain/repository/category_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	repo        repository.CategoryHandlerRepository
	Persistence *base.Persistence
}

// categorys constructor
func NewCategoryController(p *base.Persistence) *CategoryHandler {
	return &CategoryHandler{
		Persistence: p,
	}
}

// @Summary Add category
// @Description Create a new category in the database
// @Tags Category
// @Accept json
// @Produce json
// @Param category body entity.Category true "Category data"
// @Success 201 {object} response_entity.Response "Category created"
// @Failure 400 {object} response_entity.Response "Invalid category_id format, Unable to parse form data, Unable to get Category from form"
// @Failure 500 {object} response_entity.Response "Application AddCategory error"
// @Router /categories [post]
func (p *CategoryHandler) AddCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	var cat entity.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	p.repo = application.NewCategoryApplication(p.Persistence)
	newCategory, err := p.repo.AddCategory(&cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Category created.", newCategory))
}

// @Summary Get category
// @Description Get Category details by category_id
// @Tags Category
// @Accept json
// @Produce json
// @Param category_id path int true "category_id"
// @Success 200 {object} response_entity.Response "Successfully get Category"
// @Failure 400 {object} response_entity.Response "Invalid category_id GetCategory"
// @Failure 500 {object} response_entity.Response "Application GetCategory error"
// @Router /categories/{id} [get]
func (p *CategoryHandler) GetCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract category_id from the URL parameter
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category_id GetCategory", ""))

		return
	}

	// Call the service to get a single category by ID
	p.repo = application.NewCategoryApplication(p.Persistence)
	category, err := p.repo.GetCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get Categories.", category))

}

// @Summary Get category chain
// @Description Get category details including all parent categories by category_id
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "category_id"
// @Success 200 {object} response_entity.Response "Successfully get Category chain"
// @Failure 400 {object} response_entity.Response "Invalid category_id GetCategory"
// @Failure 500 {object} response_entity.Response "Application GetCategory error"
// @Router /categories/{id}/chain [get]
func (p *CategoryHandler) GetCategoryChain(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract category_id from the URL parameter
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category_id GetCategory", ""))

		return
	}

	// Call the service to get a single category by ID
	p.repo = application.NewCategoryApplication(p.Persistence)
	category, err := p.repo.GetCategoryChain(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get Categories.", category))

}

// @Summary Update a category
// @Description Update a category in the database by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "category_id"
// @Success 201 {object} response_entity.Response "category updated"
// @Failure 400 {object} response_entity.Response "Invalid category_id"
// @Failure 500 {object} response_entity.Response "Application UpdatCategory error"
// @Router /categories/{id} [put]
func (p *CategoryHandler) UpdateCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category_id Updatecategory", ""))
		return
	}

	p.repo = application.NewCategoryApplication(p.Persistence)
	cat, _ := p.repo.GetCategory(categoryID)

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responseContextData.ResponseData(response_entity.StatusFail, "invalid JSON", ""))
		return
	}

	// Log the JSON input
	log.Printf("Received JSON input for category update: %+v", cat)

	cat.CategoryID = categoryID

	newcategory, err := p.repo.UpdateCategory(cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "category updated. ", newcategory))
}

// @Summary Delete category
// @Description Delete an Category from the database by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "category_id"
// @Success 200 {object} response_entity.Response "Category deleted"
// @Failure 400 {object} response_entity.Response "Invalid category_id DeleteCategory"
// @Failure 500 {object} response_entity.Response "Application DeleteCategory error"
// @Router /categories/{id} [delete]
func (p *CategoryHandler) DeleteCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract category_id from the URL parameter
	CategoryIDStr := c.Param("id")
	CategoryID, err := strconv.ParseUint(CategoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category_id DeleteCategory", ""))
		return
	}

	// Call the service to get a single category by ID
	p.repo = application.NewCategoryApplication(p.Persistence)
	err = p.repo.DeleteCategory(CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single category
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Category deleted.", ""))
}
