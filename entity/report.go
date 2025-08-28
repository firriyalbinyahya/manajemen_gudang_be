package entity

type ReportSummary struct {
	TotalProducts int64 `json:"total_products"`
	TotalStock    int64 `json:"total_stock"`
	LowStockItems int64 `json:"low_stock_items"`
}
