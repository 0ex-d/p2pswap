package main

import (
	"fmt"
	"os"
	"p2pswap/handlers"
	"p2pswap/utils"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Branch[T int | float64] struct {
	left  []T
	right []T
}

type TreeStruct[T int | float64] struct {
	root Branch[T]
}

func refreshCacheRedis() {
	go func() {
		var count int
		last_no := 0
		for {
			if count > 3 {
				if count > last_no {

				}
				log.Printf("stopping chcks..%v", TreeStruct[float64]{root: Branch[float64]{left: []float64{1, 2}}})
				break
			}
			log.Println("checking again...")

			count++
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	port := "9911"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	host := "http://localhost"

	utils.InitRedis()

	router := gin.Default()

	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/user", handlers.GetUserDetails)
	router.POST("/trade", handlers.CreateTrade)
	// router.GET("/trade/:id", handlers.GetTrade)
	router.GET("/trades", handlers.GetAllTrades)
	refreshCacheRedis()

	log.Infof("Starting server on %s:%s", host, port)
	router.Run(fmt.Sprintf(":%s", port))
}
