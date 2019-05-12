package sqlite

import "time"

type ValueReport struct {
	ValueReportId		uint			`json:"value_report_id"`
	PrintTime 			time.Time		`json:"print_time"`
	TotalSKU 			uint			`json:"total_sku"`
	TotalStock 			uint			`json:"total_stock"`
	StockValue			[]StockValue	`json:"stock_value"`
	StockValueString	string			`json:"stock_value_string"`
}