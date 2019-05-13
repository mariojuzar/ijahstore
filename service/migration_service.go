package service

import (
	"bytes"
	"encoding/csv"
	jsoniter "github.com/json-iterator/go"
	"ijahstore/dao/sqlite"
	"ijahstore/libraries/util"
	"strconv"
)

type MigrationService interface {
	ExportToCSVReportValueById(id uint) (*bytes.Buffer, error)
	ExportToCSVReportSalesById(id uint) (*bytes.Buffer, error)
}

type migrationService struct {

}

func NewMigrationService() MigrationService  {
	return migrationService{}
}

func (migrationService) ExportToCSVReportValueById(id uint) (*bytes.Buffer, error) {
	var report sqlite.ValueReport
	var stocks []sqlite.StockValue
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	databaseService.db.Model(report).First(&report, "value_report_id = ?", id)

	_ = json.UnmarshalFromString(report.StockValueString, &stocks)

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	_ = w.Write([]string{
		"SKU",
		"Name",
		"Quantity",
		"Average Purchases",
		"Total",
	})

	for _, valStock := range stocks {
		_ = w.Write([]string{
			valStock.SKU,
			valStock.Name,
			strconv.Itoa(int(valStock.TotalStock)),
			util.PrettifyPrice("IDR", valStock.AvgPurchase),
			util.PrettifyPrice("IDR", valStock.TotalValue),
		})
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return b, err
	}

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return b, err[0]
	}

	return b, nil
}

func (migrationService) ExportToCSVReportSalesById(id uint) (*bytes.Buffer, error) {
	var report sqlite.SaleReport
	var orders []sqlite.SaleOrder
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	databaseService.db.Model(report).First(&report, "sale_report_id = ?", id)

	_ = json.UnmarshalFromString(report.SaleOrderString, &orders)

	_ = w.Write([]string{
		"Order ID",
		"Time",
		"SKU",
		"Name",
		"Quantity",
		"Sell Price",
		"Total",
		"Purchase Price",
		"Profit",
	})

	for _, sale := range orders {
		var stocks []sqlite.SaleStock

		_ = json.UnmarshalFromString(sale.SaleStockString, &stocks)

		for _, st := range stocks {
			_ = w.Write([]string{
				sale.OrderId,
				sale.Time.Format("2006-01-02 00:00:00"),
				st.SKU,
				st.Name,
				strconv.Itoa(int(st.Quantity)),
				util.PrettifyPrice("IDR", st.SellPrice),
				util.PrettifyPrice("IDR", st.TotalPrice),
				util.PrettifyPrice("IDR", st.PurchasedPrice),
				util.PrettifyPrice("IDR", st.Profit),
			})
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return b, err
	}

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return b, err[0]
	}

	return b, nil
}