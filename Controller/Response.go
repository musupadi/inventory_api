package Controller

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, data interface{}) {
	if isEmpty(data) {
		c.JSON(status, Response{
			Message: "Failed",
			Code:    1,
			Data:    []interface{}{}, // Create an empty array if data is nil or empty
		})
	} else {
		c.JSON(status, Response{
			Message: "Success",
			Code:    0,
			Data:    data,
		})
	}

}
func ResponseJSONString(c *gin.Context, status int, data string) {
	if isEmpty(data) {
		c.JSON(status, Response{
			Message: "Failed",
			Code:    1,
			Data:    "", // Create an empty array if data is nil or empty
		})
	} else {
		c.JSON(status, Response{
			Message: "Success",
			Code:    0,
			Data:    data,
		})
	}

}

// Helper function to check if the data is empty
func isEmpty(data interface{}) bool {
	if data == nil {
		return true
	}
	// Use reflection to check for empty slices and arrays
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return value.Len() == 0
	default:
		return false
	}
}
