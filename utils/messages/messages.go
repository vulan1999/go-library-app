package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessageJSON(c *gin.Context, code int, data interface{}) {
	c.IndentedJSON(code, gin.H{
		"code":    code,
		"meesage": http.StatusText(code),
		"data":    data,
	})
}
