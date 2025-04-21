package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/models"
	"github.com/vulan1999/go-library-app/utils/messages"
	queryerrors "github.com/vulan1999/go-library-app/utils/query_errors"
	"gorm.io/gorm"
)

// @BasePath /
// @Summary Get book by id
// @Tag book
// @Accept json
// @Param id path int true "Book ID"
// @Produce json
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	id := c.Param("id")

	var target models.Book

	result := config.Db.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id", "Name")
		}).
		Preload("Language", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id", "Description")
		}).
		Preload("OriginalBook", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id", "Title")
		}).
		First(&target, id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}
