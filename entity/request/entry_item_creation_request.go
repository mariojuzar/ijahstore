package request

type EntryItemCreationRequest struct {
	StockId 		uint		`json:"stock_id"`
	AmountOrder 	uint		`json:"amount_order"`
	AmountReceived	uint		`json:"amount_received"`
	PurchasePrice	uint		`json:"purchase_price"`
	ReceiptNumber	string		`json:"receipt_number"`
	Note 			string		`json:"note"`
}
