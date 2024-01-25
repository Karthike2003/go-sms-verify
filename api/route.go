package api

import "github.com/gin-gonic/gin"

type Config struct {	
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/OTP", app.sendOTP)
	app.Router.POST("/verifyOTP", app.verifyOTP)
}