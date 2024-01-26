package stock

import "gorm.io/gorm"

type StockConfig struct {
	Db *gorm.DB
}

type StockDBResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Notation    string `json:"notation"`
	StockBoard  string `json:"stock_board"`
	CompanyName string `json:"company_name"`
	IsSuspended bool   `json:"is_suspended"`
}
