package main

import (
	"github.com/gin-gonic/gin"
	"monitoring/api"
	"monitoring/config"
	"monitoring/database"
	"monitoring/handlers"
	"os"
)

func main() {
	var cfg *config.Config
	var err error
	if len(os.Args) == 2 {
		cfg, err = config.ParseConfig(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		cfg, err = config.ParseConfig("./config.yml")
		if err != nil {
			panic(err)
		}
	}

	db := database.Init(cfg)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("resources", api.GetResources)
	}

	router.GET("/", handlers.GET_Index)
	router.GET("/vpn", func(c *gin.Context) {
		handlers.GET_VPN(c, db)
	})

	err = router.Run("0.0.0.0:80")
	if err != nil {
		panic(err)
	}
}
