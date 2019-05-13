package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ijahstore/dao/sqlite"
	"ijahstore/entity/request"
)

func GetRequestBodyStockItem(c *gin.Context, item *sqlite.StockItem)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(item)
}

func GetRequestBodyStockItemUpdateRequest(c *gin.Context, item *request.StockItemUpdateRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(item)
}

func GetRequestBodyListOrder(c *gin.Context, orders *request.OrderRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(orders)
}

func GetRequestBodyEntryItemCreation(c *gin.Context, entry *request.EntryItemCreationRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(entry)
}

func GetRequestBodyEntryItemUpdate(c *gin.Context, entry *request.EntryItemUpdateRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(entry)
}

func GetRequestBodyOutcomeItem(c *gin.Context, out *request.OutComeItemCreationRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(out)
}

func GetRequestBodyReportSales(c *gin.Context, out *request.ReportSalesRequest)  {
	decoder := json.NewDecoder(c.Request.Body)
	fmt.Println(c.Request.Body)

	_ = decoder.Decode(out)
}

func GetRequestBodyReportValue(c *gin.Context, out *request.ReportValueRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(out)
}