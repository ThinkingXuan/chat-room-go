package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeSuccessJSON return success response
func MakeSuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func MakeSuccessString(c *gin.Context, data interface{}) {
	c.String(http.StatusOK, data.(string))
}

// MakeFail return fail response
func MakeFail(c *gin.Context, message interface{}) {
	c.String(http.StatusBadRequest, message.(string))
}
