package stock

import (
	"github.com/gin-gonic/gin"
	baseResponse "ijahstore/entity/response"
	"ijahstore/service"
	"net/http"
	"time"
)

func getStockService() service.StockService  {
	return service.NewStockService()
}

var srv = getStockService()

func GetAllCurrentStock(c *gin.Context)  {
	current, err := srv.GetAllCurrentStock()

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
		response.Data = current

		c.JSON(http.StatusOK, response)
	}
}