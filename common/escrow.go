package common

import (
	"errors"
	"net/http"
	"p2pswap/utils"

	"github.com/gin-gonic/gin"
)

var (
	errorLockEscrow = errors.New("failed to lock escrow")
)

type EscrowParams struct {
	TradeID string  `json:"trade_id"`
	Amount  float64 `json:"amount"`
	AssetID string  `json:"asset_id"`
}

func LockEscrow(params *EscrowParams) (interface{}, error) {
	err := utils.RedisClient.HSet(utils.Ctx, "escrow:"+params.TradeID, map[string]interface{}{
		"amount":   params.Amount,
		"asset_id": params.AssetID,
		"locked":   true,
	}).Err()

	if err != nil {
		return nil, errorLockEscrow
	}
	return nil, nil
}

func ReleaseEscrow(c *gin.Context) {
	tradeID := c.Query("trade_id")
	escrowData, err := utils.RedisClient.HGetAll(utils.Ctx, "escrow:"+tradeID).Result()
	if err != nil || len(escrowData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escrow not found"})
		return
	}
	err = utils.RedisClient.HSet(utils.Ctx, "escrow:"+tradeID, "locked", false).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to release escrow"})
		return
	}
	err = utils.RedisClient.HSet(utils.Ctx, "trade:"+tradeID, "status", "completed").Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update trade status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Escrow released successfully"})
}
