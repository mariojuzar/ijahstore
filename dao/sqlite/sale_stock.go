package sqlite

type SaleStock struct {
	SaleStockId		uint	`json:"sale_stock_id"`
	StockItem
	Quantity 		uint	`json:"quantity"`
	SellPrice		uint	`json:"sell_price"`
	TotalPrice 		uint	`json:"total_price"`
	PurchasedPrice 	uint	`json:"purchased_price"`
	Profit 			uint	`json:"profit"`
}
