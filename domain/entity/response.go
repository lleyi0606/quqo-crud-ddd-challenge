package entity

import "github.com/gin-gonic/gin"

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

type Response struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

type ServiceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseContext struct {
	Ctx *gin.Context
}

func (self *ResponseContext) ResponseData(status string, message string, data interface{}) Response {

	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return response
}
