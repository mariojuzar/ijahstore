package item

import (
	"github.com/gin-gonic/gin"
	"ijahstore/dao/sqlite"
	baseResponse "ijahstore/entity/response"
	"ijahstore/libraries/util"
	"ijahstore/service"
	"net/http"
	"time"
)

func getItemService() service.ItemService {
	return service.NewItemService()
}

var srv = getItemService()

func AddStockItem(c *gin.Context)  {
	id := srv.GenerateSKUID()

	requestBody := util.GetRequestBody(c)

	stockItem := sqlite.StockItem{
		SKUID:	id,
		SKU:	util.GenerateSKU(id, requestBody["size"], requestBody["colour"]),
		Name: 	requestBody["name"],
		Size:	requestBody["size"],
		Colour: requestBody["colour"],
	}

	err := srv.AddStockItem(&stockItem)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusCreated
		response.Message = http.StatusText(http.StatusCreated)
		response.Data = stockItem

		c.JSON(http.StatusCreated, response)
	}
}

func UpdateStockItem(c *gin.Context)  {
	requestBody := util.GetRequestBody(c)

	stockItem := sqlite.StockItem{
		ID:		util.StrToUint(requestBody["id"]),
		Name: 	requestBody["name"],
		Size:	requestBody["size"],
		Colour: requestBody["colour"],
	}

	err := srv.UpdateStockItem(&stockItem)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = stockItem

		c.JSON(http.StatusAccepted, response)
	}
}

func DeleteStockItem(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	err := srv.DeleteStockItem(id)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = nil

		c.JSON(http.StatusAccepted, response)
	}
}

func GetStockItem(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	stockItem, err := srv.GetStockItem(id)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = stockItem

		c.JSON(http.StatusAccepted, response)
	}
}

func GetAllStockItem(c *gin.Context)  {
	stockItem, err := srv.GetAllStockItem()

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = stockItem

		c.JSON(http.StatusAccepted, response)
	}
}