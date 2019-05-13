package sqlite

import "time"

type SaleReport struct {
	SaleReportId	uint		`json:"sale_report_id"`
	StartDate 		time.Time	`json:"start_date"`
	EndDate 		time.Time	`json:"end_date"`
	PrintTime 		time.Time	`json:"print_time"`
	TotalRevenue 	uint        `json:"totalRevenue"`
	TotalProfit  	uint        `json:"totalProfit"`
	TotalSale    	uint        `json:"totalSale"`
	TotalStock   	uint        `json:"totalStock"`
	SaleOrder 		[]SaleOrder	`json:"sale_order"`
	SaleOrderString string		`json:"sale_order_string"`
}
