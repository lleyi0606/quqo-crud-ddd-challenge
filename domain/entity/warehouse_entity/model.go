package warehouse_entity

type Warehouse struct {
	WarehouseID uint64  `gorm:"autoIncrement" json:"warehouse_id"`
	Name        string  `gorm:"size:100" json:"name"`
	Address     string  `gorm:"size:255" json:"address"`
	Longitude   float64 `gorm:"type:double precision" json:"longitude"`
	Latitude    float64 `gorm:"type:double precision" json:"latitude"`
}
