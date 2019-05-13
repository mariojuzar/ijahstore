package request

type OrderCreationRequest struct {
	StockId 	uint	`json:"stock_id"`
	Quantity 	uint	`json:"quantity"`
	SellPrice 	uint	`json:"sell_price"`
}
