package sqlite

import "time"

type OutcomeStockLog struct {
	OutcomeId 		uint		`json:"outcome_id"`
	Time 			time.Time	`json:"time"`
	StockItem
	OrderId 		string		`json:"order_id"`
	AmountDelivered uint		`json:"amount_delivered"`
	SellPrice		uint		`json:"sell_price"`
	TotalPrice		uint		`json:"total_price"`
	Note 			string		`json:"note"`
}