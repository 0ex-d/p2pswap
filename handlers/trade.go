package handlers

import (
	"net/http"
	"p2pswap/common"
	"p2pswap/models"
	"p2pswap/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// {
// 	"tether": {
// 	  "gbp": 0.785446
// 	}
//   }

func CreateTrade(c *gin.Context) {
	var trade models.Trade
	var err error
	if err = c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userD, err := common.GetActorById(trade.UserID)
	if err != nil {
		log.Error("user not found or authorized")
		return
	}

	log.Info("user not found or authorized", userD)
	return

	trade.ID = uuid.New().String()
	trade.Status = "open"

	var createEsc common.EscrowParams
	createEsc.Amount = trade.Amount

	log.Infof("Locking escrow Trade:%v, Esc:%v", createEsc, trade)
	common.LockEscrow(&createEsc)
	return

	err = utils.RedisClient.HSet(utils.Ctx, "trade:"+trade.ID, map[string]interface{}{
		"user_id":  trade.UserID,
		"type":     trade.Type,
		"asset_id": trade.AssetID,
		"price":    trade.Price,
		"amount":   trade.Amount,
		"status":   trade.Status,
	}).Err()

	if err != nil {
		log.Println("bad")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create trade"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Trade created successfully", "trade": trade})
}

func GetUserTrades(c *gin.Context) {
	userID := c.Query("id") // Pass userID as a query parameter

	keys, err := utils.RedisClient.Keys(utils.Ctx, "trade:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trades"})
		return
	}

	userTrades := []map[string]string{}
	for _, key := range keys {
		tradeData, _ := utils.RedisClient.HGetAll(utils.Ctx, key).Result()
		if tradeData["user_id"] == userID {
			userTrades = append(userTrades, tradeData)
		}
	}

	if len(userTrades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No trades found for user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trades": userTrades})
}

func GetAllTrades(c *gin.Context) {
	keys, err := utils.RedisClient.Keys(utils.Ctx, "trade:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trades"})
		return
	}

	trades := []map[string]string{}
	for _, key := range keys {
		tradeData, _ := utils.RedisClient.HGetAll(utils.Ctx, key).Result()
		trades = append(trades, tradeData)
	}

	c.JSON(http.StatusOK, gin.H{"trades": trades})
}
