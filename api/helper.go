package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status int 	   `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

var validate = validator.New()
func (app *Config) validateBody(c *gin.Context, body interface{}) error {
	err := c.BindJSON(&body)
	if err != nil {
		return err
	}

	err = validate.Struct(body)
	if err != nil {
		return err
	}

	return nil
} 

func (app *Config) writeJSON(c *gin.Context, status int, message string, data ...interface{}) {
	c.JSON(status, jsonResponse{
		Status: status,
		Message: "success",
		Data: data,
	})
}

func (app *Config) errorJSON(c *gin.Context, err error,status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{
		Status: statusCode,
		Message: err.Error(),
	})
}