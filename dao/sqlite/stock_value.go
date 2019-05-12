package sqlite

type StockValue struct {
	StockItem
	TotalStock 	uint 	`json:"total_stock"`
	AvgPurchase	uint	`json:"avg_purchase"`
	TotalValue	uint	`json:"total_value"`
}
