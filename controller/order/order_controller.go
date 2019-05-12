package order

import (
	"github.com/gin-gonic/gin"
	"ijahstore/entity/request"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getOrderService() service.OrderService  {
	return service.NewOrderService()
}

var srv = getOrderService()

func AddOrder(c *gin.Context)  {
	var orders request.OrderRequest
	util.GetRequestBodyListOrder(c, &orders)

	res, err := srv.AddOrder(orders.Data)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = res

		c.JSON(http.StatusOK, response)
	}
}