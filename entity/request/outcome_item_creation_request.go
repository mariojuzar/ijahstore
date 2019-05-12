package request

type OutComeItemCreationRequest struct {
	StockId 		uint	`json:"stock_id"`
	SaleOrderId	uint	`json:"sale_order_id"`
	AmountDelivered	uint	`json:"amount_delivered"`
}
