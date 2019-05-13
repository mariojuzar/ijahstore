package migration

import (
	"github.com/gin-gonic/gin"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getMigrationService() service.MigrationService {
	return service.NewMigrationService()
}

var srv = getMigrationService()

func GetCSVFromValueReport(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	b, err := srv.ExportToCSVReportValueById(id)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = nil

		date := time.Now().Local()

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+date.Format("2006-01-02")+"_value_report.csv")
		c.Data(http.StatusOK, "text/csv", b.Bytes())
	}
}

func GetCSVFromVSaleReport(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	b, err := srv.ExportToCSVReportSalesById(id)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = nil

		date := time.Now().Local()

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+date.Format("2006-01-02")+"_sales_report.csv")
		c.Data(http.StatusOK, "text/csv", b.Bytes())
	}
}