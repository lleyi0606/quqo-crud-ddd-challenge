package image

import (
	"log"
	entity "products-crud/domain/entity/image_entity"
	repository "products-crud/domain/repository/image_repository"
	base "products-crud/infrastructure/persistences"
)

type imageRepo struct {
	p *base.Persistence
}

func NewImageRepository(p *base.Persistence) repository.ImageRepository {
	return &imageRepo{p}
}

func (r imageRepo) AddImage(img *entity.Image) (*entity.Image, error) {
	log.Println("Adding new inventory ", img.ProductID, "...")

	if err := r.p.ProductDb.Debug().Create(&img).Error; err != nil {
		return nil, err
	}

	// add to search repo
	// searchRepo := search.NewSearchRepository(r.p, "algolia")
	// err := searchRepo.AddInventory(img)
	// if err != nil {
	// 	return nil, err
	// }

	log.Println(img.ProductID, " created.")
	return img, nil
}
