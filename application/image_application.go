package application

import (
	entity "products-crud/domain/entity/image_entity"
	repository "products-crud/domain/repository/image_repository"
	"products-crud/infrastructure/implementations/image"
	base "products-crud/infrastructure/persistences"
)

type imageApp struct {
	p *base.Persistence
}

func NewImageApplication(p *base.Persistence) repository.ImageHandlerRepository {
	return &imageApp{p}
}

func (u *imageApp) AddImage(img *entity.ImageInput) (*entity.Image, error) {
	repoInventory := image.NewImageRepository(u.p)
	return repoInventory.AddImage(img)
}

func (u *imageApp) GetImage(id uint64) ([]entity.Image, error) {
	repoInventory := image.NewImageRepository(u.p)
	return repoInventory.GetImage(id)
}
