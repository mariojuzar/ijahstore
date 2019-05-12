package sqlite

type StockValue struct {
	StockValueId 	uint	`json:"stock_value_id"`
	StockItem
	TotalStock 		uint 	`json:"total_stock"`
	AvgPurchase		uint	`json:"avg_purchase"`
	TotalValue		uint	`json:"total_value"`
}
