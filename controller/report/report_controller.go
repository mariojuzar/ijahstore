package report

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ijahstore/entity/request"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getNewReportService() service.ReportService  {
	return service.NewReportService()
}

var srv = getNewReportService()

func GetSalesReport(c *gin.Context)  {
	var req request.ReportSalesRequest

	util.GetRequestBodyReportSales(c, &req)

	fmt.Println(req)

	res, err := srv.GetSalesReport(req)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = res

		c.JSON(http.StatusOK, response)
	}
}

func GetValueReport(c *gin.Context)  {
	var req request.ReportValueRequest

	util.GetRequestBodyReportValue(c, &req)

	res, err := srv.GetValueReport(req)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = res

		c.JSON(http.StatusOK, response)
	}
}