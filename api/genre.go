package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/models"
	"github.com/vulan1999/go-library-app/utils/messages"
	queryerrors "github.com/vulan1999/go-library-app/utils/query_errors"
)

func GetAllGenre(c *gin.Context) {
	var genres []models.Genre

	result := config.Db.Find(&genres)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &genres)
	}
}

func GetGenreById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Genre

	result := config.Db.Take(&target, "id = ?", target_id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
