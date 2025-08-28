package entity

type Product struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(255);not null" json:"product_name"`
	SKU         string `gorm:"type:varchar(100);unique;not null" json:"sku"`
	Quantity    int    `gorm:"type:int;not null" json:"quantity"`
	Location    string `gorm:"type:varchar(255);not null" json:"location"`
	Status      string `gorm:"type:varchar(50);not null" json:"status"`
}

type CreateProductRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	SKU         string `json:"sku" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateProductRequest struct {
	ProductName *string `json:"product_name"`
	Quantity    *int    `json:"quantity,omitempty"`
	Location    *string `json:"location"`
	Status      *string `json:"status"`
}

type PaginatedProduct struct {
	Items       []Product `json:"items"`
	TotalItems  int64     `json:"total_items"`
	TotalPages  int       `json:"total_pages"`
	CurrentPage int       `json:"current_page"`
}
