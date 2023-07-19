package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    CodeType    `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseSuccess 返回状态码，信息和数据
func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	c.JSON(http.StatusOK, responseData)
}

// ResponseError 返回状态码和错误码信息
func ResponseError(c *gin.Context, code CodeType) {
	responseData := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}

// ResponseErrorWithMsg 返回错误码和信息
func ResponseErrorWithMsg(c *gin.Context, code CodeType, msg interface{}) {
	responseData := &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}
