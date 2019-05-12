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
		ItemID:    	srv.GenerateID(),
		SKUID:	id,
		SKU:	util.GenerateSKU(id, requestBody["size"], requestBody["colour"]),
		Name: 	util.PrettifyName(requestBody["name"], requestBody["size"], requestBody["colour"]),
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
		ItemID:	util.StrToUint(requestBody["id"]),
		Name: 	util.PrettifyName(requestBody["name"], requestBody["size"], requestBody["colour"]),
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
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = stockItem

		c.JSON(http.StatusOK, response)
	}
}

func GetAllStockItem(c *gin.Context)  {
	stockItem, err := srv.GetAllStockItem()

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
		response.Data = stockItem

		c.JSON(http.StatusOK, response)
	}
}