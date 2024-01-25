package handlers

import (
	"net/http"
	"products-crud/application"
	entity "products-crud/domain/entity/image_entity"
	repository "products-crud/domain/repository/image_repository"
	base "products-crud/infrastructure/persistences"

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
	var img entity.Image
	if err := c.ShouldBindJSON(&img); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	p.repo = application.NewImageApplication(p.Persistence)
	newImage, err := p.repo.AddImage(&img)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newImage)
}
