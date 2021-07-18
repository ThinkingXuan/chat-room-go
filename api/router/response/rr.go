package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeSuccess return success response
func MakeSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

//// MakeSuccessPage 分页数据处理
//func MakeSuccessPage(c *gin.Context, data interface{}, ) {
//
//	c.JSON(http.StatusOK, data)
//}

// MakeFail return fail response
func MakeFail(c *gin.Context, message interface{}) {
	c.JSON(http.StatusBadRequest, message)
}
