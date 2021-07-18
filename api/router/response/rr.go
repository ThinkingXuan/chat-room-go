package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeSuccess return success response
func MakeSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// MakeSuccessPage 分页数据处理
func MakeSuccessPage(c *gin.Context, code int, data interface{}, count int, pageIndex int, pageSize int) {

	c.JSON(http.StatusOK, gin.H{"statusCode": code, "data": data, "count": count, "index": pageIndex, "size": pageSize})
}

// MakeFail return fail response
func MakeFail(c *gin.Context, message interface{}) {
	c.JSON(http.StatusBadRequest, message)
}
