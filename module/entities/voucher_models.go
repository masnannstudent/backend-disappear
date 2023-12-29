package entities

import "time"

func (VoucherModels) TableName() string {
	return "vouchers"
}

func (VoucherClaimModels) TableName() string {
	return "voucher_claims"
}

type VoucherModels struct {
	ID          uint64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Code        string     `gorm:"column:code;type:VARCHAR(255)" json:"code"`
	Category    string     `gorm:"column:category;type:VARCHAR(255)" json:"category"`
	Description string     `gorm:"column:description;type:TEXT" json:"description"`
	Discount    uint64     `gorm:"column:discount" json:"discount"`
	StartDate   time.Time  `gorm:"column:start_date;type:TIMESTAMP" json:"start_date"`
	EndDate     time.Time  `gorm:"column:end_date;type:TIMESTAMP" json:"end_date"`
	MinPurchase uint64     `gorm:"column:min_purchase" json:"min_purchase"`
	Stock       uint64     `gorm:"column:stock" json:"stock"`
	Status      string     `gorm:"column:status;type:VARCHAR(255)" json:"status"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type VoucherClaimModels struct {
	ID        uint64         `gorm:"column:id;primaryKey" json:"id"`
	UserID    uint64         `gorm:"column:user_id" json:"user_id"`
	VoucherID uint64         `gorm:"column:voucher_id" json:"voucher_id"`
	User      *UserModels    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Voucher   *VoucherModels `gorm:"foreignKey:VoucherID" json:"voucher,omitempty"`
}
