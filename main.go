package main

import (
	"github.com/gin-gonic/gin"
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
		v1.GET(path.Item)

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