package handlers

import (
	"net/http"
	"p2pswap/utils"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	id := c.Query("id")
	userData, err := utils.RedisClient.HGetAll(utils.Ctx, "user:"+id).Result()
	if err != nil || len(userData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userData})
}
