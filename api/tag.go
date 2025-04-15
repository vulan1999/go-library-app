package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/models"
	"github.com/vulan1999/go-library-app/utils/messages"
	queryerrors "github.com/vulan1999/go-library-app/utils/query_errors"
)

func GetAllTags(c *gin.Context) {
	var tags []models.Tag

	result := config.Db.Table(fmt.Sprintf("%s.tags", os.Getenv("PG_SCHEMA"))).Find(&tags)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "There is something wrong while querying"})
		log.Panic(result.Error)
	} else {
		c.IndentedJSON(http.StatusOK, &tags)
	}
}

func GetTagById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Tag

	result := config.Db.Table(fmt.Sprintf("%s.tags", os.Getenv("PG_SCHEMA"))).First(&target, "id = ?", target_id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
