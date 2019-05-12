package request

import "time"

type ReportValueRequest struct {
	PrintTime 	time.Time	`json:"print_time"`
}