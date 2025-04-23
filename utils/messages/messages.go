package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetMessageJSON(c *gin.Context, code int, data interface{}) {
	c.IndentedJSON(code, &Response{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	})
}
