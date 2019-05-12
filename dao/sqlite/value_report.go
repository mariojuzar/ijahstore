package sqlite

import "time"

type ValueReport struct {
	PrintTime 	time.Time		`json:"print_time"`
	TotalSKU 	uint			`json:"total_sku"`
	TotalStock 	uint			`json:"total_stock"`
	StockValue	[]StockValue	`json:"stock_value"`
}