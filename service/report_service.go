package service

import (
	jsoniter "github.com/json-iterator/go"
	"ijahstore/dao/sqlite"
	"ijahstore/entity/request"
	"time"
)

type ReportService interface {
	GetSalesReport(request request.ReportSalesRequest) (sqlite.SaleReport, error)
	GetValueReport(valueRequest request.ReportValueRequest) (sqlite.ValueReport, error)
}

type reportService struct {

}

func NewReportService() ReportService {
	return reportService{}
}

func (reportService) GetSalesReport(request request.ReportSalesRequest) (sqlite.SaleReport, error) {
	var report sqlite.SaleReport
	var order []sqlite.SaleOrder
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	yearEnd, monthEnd, dayEnd := request.EndDate.Date()
	yearStart, monthStart, dayStart := request.StartDate.Date()
	yearNow, monthNow, dayNow := time.Now().Date()

	startDate := time.Date(yearStart, monthStart, dayStart, 0,0,0,0, time.UTC)
	endDate := time.Date(yearEnd, monthEnd, dayEnd, 0,0,0,0, time.UTC)
	nowDate := time.Date(yearNow, monthNow, dayNow, 0,0,0,0, time.UTC)

	databaseService.db.Model(report).Where("start_date= ? and end_date = ?", startDate, endDate).Find(&report)

	if report.SaleReportId != 0 {
		if endDate.Before(nowDate) {
			_ = json.UnmarshalFromString(report.SaleOrderString, &order)

			for i := 0; i < len(order); i++ {
				var orderSales  []sqlite.SaleStock
				_ = json.UnmarshalFromString(order[i].SaleStockString, & orderSales)
				order[i].SaleStock = orderSales
			}

			report.SaleOrder = order

			return report, nil
		}
	}

	databaseService.db.Model(order).Where("time > ? and time < ?", startDate, endDate).Find(&order)

	var totalRevenue uint
	var totalProfit uint
	var totalStock uint

	for i := 0; i < len(order); i++ {
		var orderSales  []sqlite.SaleStock
		_ = json.UnmarshalFromString(order[i].SaleStockString, & orderSales)
		order[i].SaleStock = orderSales

		for _, s := range orderSales {
			totalRevenue += s.TotalPrice
			totalProfit += s.Profit
			totalStock += s.Quantity
		}
	}

	report = sqlite.SaleReport{
		SaleReportId:	generateSaleReportId(),
		StartDate:		startDate,
		EndDate: 		endDate,
		PrintTime:		time.Now(),
		TotalRevenue: 	totalRevenue,
		TotalProfit: 	totalProfit,
		TotalSale: 		uint(len(order)),
		TotalStock: 	totalStock,
		SaleOrder: 		order,
		SaleOrderString:generateSaleReportOrderString(order),
	}

	databaseService.db.Create(report)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return report, err[0]
	}

	return report, nil
}

func (reportService) GetValueReport(valueRequest request.ReportValueRequest) (sqlite.ValueReport, error) {
	var report sqlite.ValueReport
	var value []sqlite.StockValue
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	year, month, day := valueRequest.PrintTime.Date()
	printDate := time.Date(year, month, day, 0,0,0,0, time.UTC)

	databaseService.db.Model(report).Where("print_time = ?", printDate).Find(&report)

	if report.ValueReportId != 0 {
		_ = json.UnmarshalFromString(report.StockValueString, &value)

		report.StockValue = value

		return report, nil
	}

	var stocks []sqlite.StockItem
	var totalStock uint

	databaseService.db.Model(stocks).Find(&stocks)

	for _, item := range stocks {
		var current sqlite.CurrentStockItem
		var entries []sqlite.EntryStockLog
		var avg uint

		databaseService.db.Model(current).First(&current, "item_id = ?", item.ItemID)
		databaseService.db.Find(&entries, "item_id =?", item.ItemID)

		avg = uint(getPurchasePrice(entries))
		totalStock += current.CurrentStock

		val := sqlite.StockValue{
			StockValueId:	generateStockValueId(),
			StockItem:		item,
			TotalStock:		current.CurrentStock,
			AvgPurchase: 	avg,
			TotalValue:		avg * current.CurrentStock,
		}

		value = append(value, val)
	}

	valString, _ := json.MarshalToString(value)

	report = sqlite.ValueReport{
		ValueReportId: 	generateValueReportId(),
		PrintTime: 		printDate,
		TotalSKU:		uint(len(value)),
		TotalStock:		totalStock,
		StockValue: 	value,
		StockValueString:valString,
	}

	databaseService.db.Create(report)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return report, err[0]
	}

	return report, nil
}

func generateSaleReportId() uint  {
 	var sale sqlite.SaleReport

 	row, _ := databaseService.db.Model(sale).Select("MAX(sale_report_id) as sale_report_id").Find(&sale).Rows()

	_ = databaseService.db.ScanRows(row, &sale)

 	return sale.SaleReportId + 1
}

func generateValueReportId() uint  {
	var sale sqlite.ValueReport

	row, _ := databaseService.db.Model(sale).Select("MAX(value_report_id) as value_report_id").Find(&sale).Rows()

	_ = databaseService.db.ScanRows(row, &sale)

	return sale.ValueReportId + 1
}

func generateStockValueId() uint  {
	var sale sqlite.StockValue

	row, _ := databaseService.db.Model(sale).Select("MAX(stock_value_id) as stock_value_id").Find(&sale).Rows()

	_ = databaseService.db.ScanRows(row, &sale)

	return sale.StockValueId + 1
}

func generateSaleReportOrderString(order []sqlite.SaleOrder) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	res, _ := json.MarshalToString(order)

	return res
}