package sqlite

import "time"

type EntryStockLog struct {
	Time 			time.Time	`json:"time"`
	StockItem
	AmountOrder 	uint		`json:"amount_order"`
	AmountReceived	uint		`json:"amount_received"`
	PurchasePrice	uint		`json:"purchase_price"`
	TotalPrice		uint		`json:"total_price"`
	ReceiptNumber	string		`json:"receipt_number"`
	Note 			string		`json:"note"`
}