package image

import (
	"errors"
	"fmt"
	"log"
	entity "products-crud/domain/entity/image_entity"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/image_repository"
	"products-crud/infrastructure/implementations/cache"
	base "products-crud/infrastructure/persistences"

	storage_go "github.com/supabase-community/storage-go"
	"gorm.io/gorm"
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

	err = r.p.ProductDb.Debug().Where("image_id = ?", image.ImageID).Updates(&image).Error

	log.Println(img.ProductID, " created. ", res.Key)
	return &image, nil
}

func (r imageRepo) GetImage(id uint64) ([]entity.Image, error) {
	var img []entity.Image

	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	_ = cacheRepo.GetKey(fmt.Sprintf("%s%d", redis_entity.RedisImageData, id), &img)

	if img == nil {
		err := r.p.ProductDb.Debug().Where("product_id = ?", id).Find(&img).Error
		// err := r.p.ProductDb.Debug().Where("product_id = ?", id).Take(&img).Error
		if err != nil {
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		_ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisImageData, id), img, redis_entity.RedisExpirationGlobal)
	}

	return img, nil
}

func (r imageRepo) DeleteImage(id uint64) error {
	var img entity.Image

	err := r.p.ProductDb.Debug().Where("product_id = ?", id).Delete(&img).Error
	if err != nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("product not found")
	}

	// delete from CDN database
	_, err = r.p.ImageSupabaseDB.RemoveFile("images", []string{"image/storage/" + fmt.Sprint(id)})
	if err != nil {
		return err
	}

	// update cache
	// cacheRepo := cache.NewCacheRepository(r.p, "redis")
	// err = cacheRepo.DeleteRecord(fmt.Sprintf("%s%d", redis_entity.RedisImageData, id))
	// if err != nil {
	// 	return err
	// }

	return nil
}
