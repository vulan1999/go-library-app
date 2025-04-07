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

func GetAllAuthor(c *gin.Context) {
	var authors []models.Author

	result := config.Db.Table(fmt.Sprintf("%s.authors", os.Getenv("PG_SCHEMA"))).Find(&authors)

	if result.Error != nil {
		messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &authors)
	}
}

func GetAuthorById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Author
	result := config.Db.Table(fmt.Sprintf("%s.authors", os.Getenv("PG_SCHEMA"))).First(&target, "id = ?", target_id)
	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
