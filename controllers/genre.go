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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			log.Panic(result.Error)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "There is something wrong whilee querying"})
			log.Panic(result.Error)
		}
	} else {
		c.IndentedJSON(http.StatusOK, &target)
	}
}
