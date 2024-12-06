package common

import (
	"errors"
	"p2pswap/utils"
)

var (
	errorActorNotFound = errors.New("failed to get user at this time. try again")
)

func GetActorById(id string) (interface{}, error) {
	userData, err := utils.RedisClient.HGetAll(utils.Ctx, "user:"+id).Result()
	if err != nil || len(userData) == 0 {
		return nil, errorActorNotFound
	}
	return nil, nil
}
