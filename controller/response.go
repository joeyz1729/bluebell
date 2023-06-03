package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    CodeType    `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	c.JSON(http.StatusOK, responseData)
}

func ResponseError(c *gin.Context, code CodeType) {
	responseData := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}

func ResponseErrorWithMsg(c *gin.Context, code CodeType, msg interface{}) {
	responseData := &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}
