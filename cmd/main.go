package main

import (
	"github.com/Karthike2003/go-sms-verify-yt/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	api.Routes(r)

	router.Run(":8080")
}