package outcome

import (
	"github.com/gin-gonic/gin"
	"ijahstore/entity/request"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getOutcomeItemService() service.OutcomeItemService {
	return service.NewOutcomeItemService()
}

var srv = getOutcomeItemService()

func AddOutcomeItem(c *gin.Context)  {
	var out request.OutComeItemCreationRequest

	util.GetRequestBodyOutcomeItem(c, &out)

	res, err := srv.AddOutcomeItem(out)

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
		response.Data = res

		c.JSON(http.StatusOK, response)
	}
}

func GetAllOutcomeItem(c *gin.Context)  {
	res, err := srv.GetAllOutcomeItem()

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