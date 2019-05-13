package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ijahstore/entity/request"
	"io/ioutil"
)

func GetRequestBody(c *gin.Context) map[string]string {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	requestBody := make(map[string]string)
	_ = json.Unmarshal(bodyBytes, &requestBody)

	return requestBody
}

func GetRequestBodyListOrder(c *gin.Context, orders *request.OrderRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(orders)
}

func GetRequestBodyEntryItemCreation(c *gin.Context, entry *request.EntryItemCreationRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(entry)
}

func GetRequestBodyEntryItemUpdate(c *gin.Context, entry *request.EntryItemUpdateRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(entry)
}

func GetRequestBodyOutcomeItem(c *gin.Context, out *request.OutComeItemCreationRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(out)
}

func GetRequestBodyReportSales(c *gin.Context, out *request.ReportSalesRequest)  {
	decoder := json.NewDecoder(c.Request.Body)
	fmt.Println(c.Request.Body)

	_ = decoder.Decode(out)
}

func GetRequestBodyReportValue(c *gin.Context, out *request.ReportValueRequest)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(out)
}