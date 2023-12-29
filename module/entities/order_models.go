package entities

import "time"

type OrderModels struct {
	ID                    string               `gorm:"column:id;type:VARCHAR(255);primaryKey" json:"id"`
	IdOrder               string               `gorm:"column:id_order;type:VARCHAR(255)" json:"id_order"`
	AddressID             uint64               `gorm:"column:address_id" json:"address_id"`
	UserID                uint64               `gorm:"column:user_id" json:"user_id"`
	VoucherID             *uint64              `gorm:"column:voucher_id" json:"voucher_id"`
	Note                  string               `gorm:"column:note;type:VARCHAR(255)" json:"note"`
	GrandTotalGramPlastic uint64               `gorm:"column:grand_total_gram_plastic" json:"grand_total_gram_plastic"`
	GrandTotalExp         uint64               `gorm:"column:grand_total_exp" json:"grand_total_exp"`
	GrandTotalQuantity    uint64               `gorm:"column:grand_total_quantity" json:"grand_total_quantity"`
	GrandTotalPrice       uint64               `gorm:"column:grand_total_price" json:"grand_total_price"`
	ShipmentFee           uint64               `gorm:"column:shipment_fee" json:"shipment_fee"`
	AdminFees             uint64               `gorm:"column:admin_fees" json:"admin_fees"`
	GrandTotalDiscount    uint64               `gorm:"column:grand_total_discount" json:"grand_total_discount"`
	TotalAmountPaid       uint64               `gorm:"column:total_amount_paid" json:"total_amount_paid"`
	OrderStatus           string               `gorm:"column:order_status;type:VARCHAR(255)" json:"order_status"`
	PaymentStatus         string               `gorm:"column:payment_status;type:VARCHAR(255)" json:"payment_status"`
	PaymentMethod         string               `gorm:"column:payment_method;type:VARCHAR(255)" json:"payment_method"`
	ExtraInfo             string               `gorm:"column:extra_info;type:VARCHAR(255)" json:"extra_info"`
	StatusOrderDate       time.Time            `gorm:"column:status_order_date;type:timestamp" json:"status_order_date"`
	CreatedAt             time.Time            `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time            `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt             *time.Time           `gorm:"column:deleted_at;index" json:"deleted_at"`
	Address               AddressModels        `gorm:"foreignKey:AddressID" json:"address"`
	User                  UserModels           `gorm:"foreignKey:UserID" json:"user"`
	Voucher               VoucherModels        `gorm:"foreignKey:VoucherID" json:"voucher"`
	OrderDetails          []OrderDetailsModels `gorm:"foreignKey:OrderID" json:"order_details"`
}

type OrderDetailsModels struct {
	ID               uint64        `gorm:"column:id;primaryKey" json:"id"`
	OrderID          string        `gorm:"column:order_id;type:VARCHAR(255)" json:"order_id"`
	ProductID        uint64        `gorm:"column:product_id" json:"product_id"`
	Quantity         uint64        `gorm:"column:quantity" json:"quantity"`
	TotalDiscount    uint64        `gorm:"column:total_discount" json:"total_discount"`
	TotalGramPlastic uint64        `gorm:"column:total_gram_plastic" json:"total_gram_plastic"`
	TotalExp         uint64        `gorm:"column:total_exp" json:"total_exp"`
	TotalPrice       uint64        `gorm:"column:total_price" json:"total_price"`
	Product          ProductModels `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (OrderModels) TableName() string {
	return "orders"
}

func (OrderDetailsModels) TableName() string {
	return "order_details"
}
