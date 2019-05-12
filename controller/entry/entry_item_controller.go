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
	requestBody := util.GetRequestBody(c)

	entry := &request.EntryItemCreationRequest{
		AmountOrder: 	util.StrToUint(requestBody["amountOrder"]),
		AmountReceived: util.StrToUint(requestBody["amountReceived"]),
		StockId: 		util.StrToUint(requestBody["stockId"]),
		PurchasePrice: 	util.StrToUint(requestBody["purchasePrice"]),
		ReceiptNumber: 	requestBody["receiptNumber"],
		Note: 			requestBody["note"],
	}

	res, err := srv.AddEntryItem(entry)

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
	requestBody := util.GetRequestBody(c)

	entry := &request.EntryItemUpdateRequest{
		ID: 			util.StrToUint(requestBody["id"]),
		AmountReceived:	util.StrToUint(requestBody["amountReceived"]),
		Note: 			requestBody["note"],
	}

	res, err := srv.UpdateEntryItem(entry)

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