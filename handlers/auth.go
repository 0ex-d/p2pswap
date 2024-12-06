package handlers

import (
	"net/http"
	"p2pswap/models"
	"p2pswap/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	authGenericError = "user not found or Invalid username or password"
)

func LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": authGenericError,
			"Data":  nil,
		})
		return
	}

	// Search for the user in Redis
	keys, err := utils.RedisClient.Keys(utils.Ctx, "user:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": authGenericError,
			"Data":  nil,
		})
		return
	}

	for _, key := range keys {
		userData, _ := utils.RedisClient.HGetAll(utils.Ctx, key).Result()
		if userData["username"] == credentials.Username && userData["password"] == credentials.Password {
			c.JSON(http.StatusOK, gin.H{
				"Err": nil,
				"Data": gin.H{
					"ID": key[len("user:"):],
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"Error": authGenericError,
		"Data":  nil,
	})
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": authGenericError,
			"Data":  nil,
		})
		return
	}
	user.ID = uuid.New().String()
	time.Sleep(2 * time.Second)
	log.Println("Enabling KYC...")
	// Mock KYC verification
	user.KYCVerified = true

	err := utils.RedisClient.HSet(utils.Ctx, "user:"+user.ID, map[string]interface{}{
		"username":     user.Username,
		"password":     user.Password,
		"kyc_verified": user.KYCVerified,
	}).Err()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": authGenericError,
			"Data":  nil,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Err":  nil,
		"Data": user,
	})
}
