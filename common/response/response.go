package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message string
	Gin     *gin.Context
	Data    interface{}
	Token   string
}

func HttpResponse(param ParamHTTPResp) {
	if param.Err != nil {
		param.Gin.JSON(param.Code, Response{
			Status:  "error",
			Message: param.Message,
			Data:    param.Data,
		})
	} else {
		param.Gin.JSON(param.Code, Response{
			Status:  "success",
			Message: param.Message,
			Data:    param.Data,
		})
	}

}
