package entity

type Product struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name        string  `gorm:"size:100;not null;" json:"name"`
	Description *string `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;not null;" json:"price"`
	Category    string  `gorm:"size:100;not null;" json:"category"`
	Stock       int     `gorm:"type:numeric;not null;" json:"stock"`
	Image       string  `gorm:"size:255;" json:"image"`
}

type ProductToReceive struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
}

type ProductUpdate struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
}

// func addProduct(db *gorm.DB, id int, name string, desc string, price float64, cat string, stock int, image string) error {
// 	log.Println("Adding new product ", name, "...")

// 	if err := db.Create(&Product{ID: uint64(id), Name: name, Description: &desc, Price: price, Category: cat, Stock: stock, Image: image}).Error; err != nil {
// 		return err
// 	}

// 	log.Println(name, " created.")
// 	return nil
// }
