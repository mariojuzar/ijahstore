package sqlite

import "time"

type SaleReport struct {
	SaleReportId	uint		`json:"sale_report_id"`
	StartDate 		time.Time	`json:"start_date"`
	EndDate 		time.Time	`json:"end_date"`
	TotalRevenue 	uint        `json:"totalRevenue"`
	TotalProfit  	uint        `json:"totalProfit"`
	TotalSale    	uint        `json:"totalSale"`
	TotalStock   	uint        `json:"totalStock"`
	SaleStock 		[]SaleStock	`json:"sale_stock"`
}
