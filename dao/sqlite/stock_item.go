package sqlite

type StockItem struct {
	ID 				uint	`json:"id"`
	SKUID  			uint	`json:"skuid"`
	SKU 			string	`json:"sku"`
	Name 			string	`json:"name"`
	Size 			string	`json:"size"`
	Colour 			string	`json:"colour"`
}