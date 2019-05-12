package request

type OrderRequest struct {
	Data 	[]OrderCreationRequest	`json:"data"`
}
