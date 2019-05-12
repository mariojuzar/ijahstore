package sqlite

import "time"

type SaleOrder struct {
	SaleOrderId	uint		`json:"sale_order_id"`
	OrderId 	string		`json:"order_id"`
	Time 		time.Time	`json:"time"`
	SaleStock 	[]SaleStock	`json:"sale_stock"`
}
