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

func GetAllGenre(c *gin.Context) {
	var genres []models.Genre

	result := config.Db.Table(fmt.Sprintf("%s.genres", os.Getenv("PG_SCHEMA"))).Find(&genres)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "There is something wrong while querying"})
		log.Panic(result.Error)
	} else {
		c.IndentedJSON(http.StatusOK, &genres)
	}
}

func GetGenreById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Genre

	result := config.Db.Table(fmt.Sprintf("%s.genres", os.Getenv("PG_SCHEMA"))).Take(&target, "id = ?", target_id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
