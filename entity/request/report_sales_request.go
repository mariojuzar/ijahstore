package request

import "time"

type ReportSalesRequest struct {
	StartDate 	time.Time	`json:"start_date"`
	EndDate 	time.Time	`json:"end_date"`
}