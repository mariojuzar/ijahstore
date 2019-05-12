package sqlite

type SaleStock struct {
	StockItem
	Quantity 		uint	`json:"quantity"`
	SellPrice		uint	`json:"sell_price"`
	TotalPrice 		uint	`json:"total_price"`
	PurchasedPrice 	uint	`json:"purchased_price"`
	Profit 			uint	`json:"profit"`
}
