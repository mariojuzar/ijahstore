package request

type StockItemUpdateRequest struct {
	ID 		uint	`json:"id"`
	Name 	string	`json:"name"`
	Size 	string	`json:"size"`
	Colour 	string	`json:"colour"`
}
