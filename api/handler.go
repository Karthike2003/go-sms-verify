package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Karthike2003/go-sms-verify-yt/data"
)
func (app *Config) sendOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validatiteBody(c, &payload)

		newData := data.OTP{
			PhoneNumber: payload.PhoneNumber,
		}

		_,err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

	app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
}

}

func (app *Config) verifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validatiteBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		//check if code is valid
		_,err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusOK, "OTP verified successfully")
	}

}