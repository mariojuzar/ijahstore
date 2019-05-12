package sqlite

type CurrentStockItem struct {
	StockItem
	CurrentStock	uint	`json:"current_stock"`
}