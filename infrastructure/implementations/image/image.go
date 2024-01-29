package image

import (
	"fmt"
	"log"
	entity "products-crud/domain/entity/image_entity"
	repository "products-crud/domain/repository/image_repository"
	base "products-crud/infrastructure/persistences"

	storage_go "github.com/supabase-community/storage-go"
)

type imageRepo struct {
	p *base.Persistence
}

func NewImageRepository(p *base.Persistence) repository.ImageRepository {
	return &imageRepo{p}
}

func (r imageRepo) AddImage(img *entity.ImageInput) (*entity.Image, error) {
	log.Println("Adding new image ", img.ProductID, "...")

	var image entity.Image
	image.Caption = img.Caption
	image.ProductID = img.ProductID

	if err := r.p.ProductDb.Debug().Create(&image).Error; err != nil {
		return nil, err
	}

	// Open the uploaded file
	file, err := img.ImageFile.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// contentTypePNG := "image/png"
	contentTypeJPEG := "image/jpeg"

	// Add to Supabase with file options
	res, err := r.p.ImageSupabaseDB.UploadFile("images", "image/storage/"+fmt.Sprint(image.ImageID), file,
		storage_go.FileOptions{
			ContentType: &contentTypeJPEG,
		},
	)
	if err != nil {
		return nil, err
	}

	result := r.p.ImageSupabaseDB.GetPublicUrl("images", "image/storage/"+fmt.Sprint(img.ProductID))
	image.Url = result.SignedURL

	log.Println(img.ProductID, " created. ", res.Key)
	return &image, nil
}
