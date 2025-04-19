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

func GetAllTags(c *gin.Context) {
	var tags []models.Tag

	result := config.Db.Find(&tags)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &tags)
	}
}

func GetTagById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Tag

	result := config.Db.First(&target, "id = ?", target_id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
