package queryerrors

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/utils/messages"
	"gorm.io/gorm"
)

func GetQueryErrorMessage(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		messages.GetMessageJSON(c, http.StatusNotFound, nil)
	} else if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		messages.GetMessageJSON(c, http.StatusBadRequest, nil)
	} else if errors.Is(err, gorm.ErrModelValueRequired) {
		messages.GetMessageJSON(c, http.StatusBadRequest, nil)
	} else {
		messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
	}
	log.Panic(err)
}
