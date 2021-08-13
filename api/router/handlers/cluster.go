package handlers

import (
	"chat-room-go/api/router/response"
	"github.com/gin-gonic/gin"
)

func UpdateCluster(c *gin.Context) {
	response.MakeSuccessString(c, "success")
}

func CheckCluster(c *gin.Context) {
	response.MakeSuccessString(c, "success")
}
