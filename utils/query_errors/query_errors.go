package queryerrors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/utils/messages"
)

func GetQueryErrorMessage(c *gin.Context, err error) {
	messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
	log.Panic(err)
}
