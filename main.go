package main

import (
	"github.com/gin-gonic/gin"
	entryController "ijahstore/controller/entry"
	itemController "ijahstore/controller/item"
	orderController "ijahstore/controller/order"
	outcomeController "ijahstore/controller/outcome"
	reportController "ijahstore/controller/report"
	stockController "ijahstore/controller/stock"
	"ijahstore/entity/path"
	"ijahstore/entity/response"
	"net/http"
	"time"
)

func main()  {
	var baseURL = "/api/v1"
	engine := gin.Default()
	engine.RedirectTrailingSlash = false

	v1 := engine.Group(baseURL)
	{
		// item controller path
		v1.GET(path.Item, itemController.GetAllStockItem)
		v1.GET(path.ItemById, itemController.GetStockItem)
		v1.POST(path.Item, itemController.AddStockItem)
		v1.PUT(path.Item, itemController.UpdateStockItem)

		// current stock controller path
		v1.GET(path.Stock, stockController.GetAllCurrentStock)

		// entry item controller path
		v1.GET(path.EntryItem, entryController.GetAllEntryItem)
		v1.POST(path.EntryItem, entryController.AddEntryItem)
		v1.PUT(path.EntryItem, entryController.UpdateEntryItem)

		// outcome item controller path
		v1.GET(path.OutcomeItem, outcomeController.GetAllOutcomeItem)
		v1.POST(path.OutcomeItem, outcomeController.AddOutcomeItem)

		// order controller path
		v1.POST(path.Order, orderController.AddOrder)

		// report controller path
		v1.GET(path.ReportSales, reportController.GetSalesReport)
		v1.GET(path.ReportValue, reportController.GetValueReport)

		// migration controller path
		v1.POST(path.MigrationImport)
		v1.GET(path.MigrationExport)
	}

	engine.NoRoute(func(context *gin.Context) {
		var resp = &response.BaseResponse{
			ServerTime:	time.Now(),
		}

		resp.Code = http.StatusNotFound
		resp.Message = "Route not found"

		context.JSON(http.StatusNotFound, resp)
	})

	_ = engine.Run(":7090")
}