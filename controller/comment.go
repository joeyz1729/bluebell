package controller

import "github.com/gin-gonic/gin"

func TODO(c *gin.Context) {
	ResponseError(c, CodeServerBusy)
	return
}
