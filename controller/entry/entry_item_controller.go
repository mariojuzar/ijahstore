package entry

import (
	"github.com/gin-gonic/gin"
	"ijahstore/entity/request"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getEntryItemService() service.EntryItemService  {
	return service.NewEntryItemService()
}

var srv = getEntryItemService()

func AddEntryItem(c *gin.Context)  {
	var entry request.EntryItemCreationRequest
	util.GetRequestBodyEntryItemCreation(c, &entry)

	res, err := srv.AddEntryItem(&entry)

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

func UpdateEntryItem(c *gin.Context)  {
	var entry request.EntryItemUpdateRequest
	util.GetRequestBodyEntryItemUpdate(c, &entry)

	res, err := srv.UpdateEntryItem(&entry)

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

func GetAllEntryItem(c *gin.Context)  {
	entry, err := srv.GetAllEntryItem()

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
		response.Data = entry

		c.JSON(http.StatusOK, response)
	}
}