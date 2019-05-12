package sqlite

import "time"

type SaleOrder struct {
	OrderId 	string		`json:"order_id"`
	Time 		time.Time	`json:"time"`
	SaleStock 	[]SaleStock	`json:"sale_stock"`
}
