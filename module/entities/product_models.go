package entities

import (
	"time"
)

type ProductModels struct {
	ID            uint64                `gorm:"column:id;primaryKey" json:"id"`
	Name          string                `gorm:"column:name;type:varchar(255)" json:"name"`
	Description   string                `gorm:"column:description;type:text" json:"description"`
	GramPlastic   uint64                `gorm:"column:gram_plastic;type:bigint" json:"gram_plastic"`
	Price         uint64                `gorm:"column:price;type:bigint" json:"price"`
	Stock         uint64                `gorm:"column:stock;type:bigint" json:"stock"`
	Discount      uint64                `gorm:"column:discount;type:bigint" json:"discount"`
	Exp           uint64                `gorm:"column:exp;type:bigint" json:"exp"`
	Rating        float64               `gorm:"column:rating;type:DECIMAL(5,2)" json:"rating"`
	TotalReview   uint64                `gorm:"column:total_review;type:bigint" json:"total_review"`
	CreatedAt     time.Time             `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time             `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time            `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	ProductPhotos []ProductPhotosModels `gorm:"foreignKey:ProductID" json:"product_photos"`
	ProductReview []ReviewModels        `gorm:"foreignKey:ProductID" json:"review"`
	Categories    []CategoryModels      `gorm:"many2many:product_categories;" json:"categories"`
}

type ProductPhotosModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	ProductID uint64     `gorm:"column:product_id" json:"product_id"`
	ImageURL  string     `gorm:"column:url;type:varchar(255)" json:"url"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (ProductModels) TableName() string {
	return "products"
}

func (ProductPhotosModels) TableName() string {
	return "product_photos"
}
