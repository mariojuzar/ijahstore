package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
