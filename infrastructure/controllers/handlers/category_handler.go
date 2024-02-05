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

// @Summary Add Category
// @Description Add an Category to the database
// @Tags Category
// @Accept mpfd
// @Produce json
// @Param category_id formData int64 true "category ID"
// @Param caption formData string false "Caption"
// @Param Category_file formData file true "Category file"
// @Success 201 {object} response_entity.Response "Category created"
// @Failure 400 {object} response_entity.Response "Invalid category ID format, Unable to parse form data, Unable to get Category from form"
// @Failure 500 {object} response_entity.Response "Application AddCategory error"
// @Router /Categorys [post]
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

// @Summary Get Categorys
// @Description Get Category details by category ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "category ID"
// @Success 200 {object} response_entity.Response "Successfully get Categorys"
// @Failure 400 {object} response_entity.Response "Invalid category ID GetCategory"
// @Failure 500 {object} response_entity.Response "Application GetCategory error"
// @Router /Categorys/{id} [get]
func (p *CategoryHandler) GetCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract category ID from the URL parameter
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category ID GetCategory", ""))

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

// @Summary Update a
// @Description Update a category in the database by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "category ID"
// @Success 201 {object} response_entity.Response "category updated"
// @Failure 400 {object} response_entity.Response "Invalid category ID"
// @Failure 500 {object} response_entity.Response "Application Updatecategory error"
// @Router /categorys/{id} [put]
func (p *CategoryHandler) UpdateCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category ID Updatecategory", ""))
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

// @Summary Delete Category
// @Description Delete an Category from the database by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response_entity.Response "Category deleted"
// @Failure 400 {object} response_entity.Response "Invalid Category ID DeleteCategory"
// @Failure 500 {object} response_entity.Response "Application DeleteCategory error"
// @Router /Categorys/{id} [delete]
func (p *CategoryHandler) DeleteCategory(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract category ID from the URL parameter
	CategoryIDStr := c.Param("id")
	CategoryID, err := strconv.ParseUint(CategoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid category ID DeleteCategory", ""))
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
