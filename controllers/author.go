package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/models"
	"gorm.io/gorm"
)

func GetAllAuthor(c *gin.Context) {
	var authors []models.Author

	result := config.Db.Table(fmt.Sprintf("%s.authors", os.Getenv("PG_SCHEMA"))).Find(&authors)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "There is something wrong while querying"})
		log.Panic(result.Error)
	} else {
		c.IndentedJSON(http.StatusOK, &authors)
	}
}

func GetAuthorById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Author
	result := config.Db.Table(fmt.Sprintf("%s.authors", os.Getenv("PG_SCHEMA"))).Take(&target, "id = ?", target_id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no such id"})
			log.Panic(result.Error)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "There is something wrong while querying"})
			log.Panic(result.Error)
		}
	} else {
		c.IndentedJSON(http.StatusOK, &target)
	}
}
