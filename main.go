package main

import (
	"github.com/gin-gonic/gin"
	itemController "ijahstore/controller/item"
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
		v1.GET(path.EntryItem)

		// outcome item controller path
		v1.GET(path.OutcomeItem)

		// report controller path
		v1.GET(path.Report)

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