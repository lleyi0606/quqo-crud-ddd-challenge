package handlers

import (
	"log"
	"net/http"
	"products-crud/application"
	response_entity "products-crud/domain/entity"
	entity "products-crud/domain/entity/image_entity"
	repository "products-crud/domain/repository/image_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	repo        repository.ImageHandlerRepository
	Persistence *base.Persistence
}

// Products constructor
func NewImageController(p *base.Persistence) *ImageHandler {
	return &ImageHandler{
		Persistence: p,
	}
}

func (p *ImageHandler) AddImage(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Parse the form data, including the uploaded file
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Unable to parse form data.", ""))
		return
	}

	// Retrieve the text fields from the form data
	productIdStr := c.Request.FormValue("product_id")
	caption := c.Request.FormValue("caption")

	productId, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		log.Print(productIdStr)

		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid productId format", ""))
		return
	}

	file, fileHeader, err := c.Request.FormFile("image_file")
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responseContextData.ResponseData(response_entity.StatusFail, "Unable to get image from form", ""))

		return
	}
	defer file.Close()

	// Create an entity.ImageInput with the received data
	img := entity.ImageInput{
		ProductID: productId,
		Caption:   caption,
		ImageFile: fileHeader,
	}

	p.repo = application.NewImageApplication(p.Persistence)
	newImage, err := p.repo.AddImage(&img)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))

		return
	}
	c.JSON(http.StatusCreated, responseContextData.ResponseData(response_entity.StatusSuccess, "Image created. ", newImage))

}

func (p *ImageHandler) GetImage(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID GetInventory", ""))

		return
	}

	// Call the service to get a single product by ID
	p.repo = application.NewImageApplication(p.Persistence)
	product, err := p.repo.GetImage(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Successfully get images.", product))

}

func (p *ImageHandler) DeleteImage(c *gin.Context) {
	responseContextData := response_entity.ResponseContext{Ctx: c}

	// Extract product ID from the URL parameter
	imageIDStr := c.Param("id")
	imageID, err := strconv.ParseUint(imageIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(response_entity.StatusFail, "Invalid product ID DeleteImage", ""))
		return
	}

	// Call the service to get a single product by ID
	p.repo = application.NewImageApplication(p.Persistence)
	err = p.repo.DeleteImage(imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseContextData.ResponseData(response_entity.StatusFail, err.Error(), ""))
		return
	}

	// Respond with the single product
	c.JSON(http.StatusOK, responseContextData.ResponseData(response_entity.StatusSuccess, "Image deleted.", ""))
}
