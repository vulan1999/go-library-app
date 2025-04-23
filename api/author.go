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

// @Summary Get All Authors
// @Tags Authors
// @ID get-all-authors
// @Accept json
// @Produce json
// @Success 200 {object} messages.Response{data=[]models.Author}
// @Failure 500 {object} messages.Response{data=nil}
// @Router /authors/ [get]
func GetAllAuthor(c *gin.Context) {
	var authors []models.Author

	result := config.Db.Find(&authors)

	if result.Error != nil {
		messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &authors)
	}
}

// @Summary Get Author By Id
// @Tags Authors
// @ID get-author-by-id
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} messages.Response{data=models.Author}
// @Failure 500 {object} messages.Response{data=nil}
// @Router /authors/{id} [get]
func GetAuthorById(c *gin.Context) {
	target_id := c.Param("id")
	var target models.Author
	result := config.Db.First(&target, "id = ?", target_id)
	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}

// @Summary Create Author
// @Tags Authors
// @ID create-author
// @Accept json
// @Produce json
// @Param body body models.AuthorCreateRequest true "Author detail to create"
// @Success 201 {object} messages.Response{data=models.Author}
// @Failure 500 {object} messages.Response{data=nil}
// @Router /authors [post]
func CreaeAuthor(c *gin.Context) {
	var body models.AuthorCreateRequest

	if err := c.BindJSON(&body); err != nil {
		messages.GetMessageJSON(c, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	newAuthor := models.Author{Name: body.Name}

	create_result := config.Db.Create(&newAuthor)

	if create_result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, create_result.Error)
		return
	}
	messages.GetMessageJSON(c, http.StatusCreated, &newAuthor)
}

// @Summary Update Author
// @Tags Authors
// @ID update-author-by-id
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} messages.Response{data=models.Author}
// @Failure 500 {object} messages.Response{data=nil}
// @Router /authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	update_id := c.Param("id")
	var update_author models.Author
	c.ShouldBindJSON(&update_author)
	var target_author models.Author
	check_result := config.Db.First(&target_author, "id = ?", update_id)
	if check_result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, check_result.Error)
	} else {
		config.Db.Model(&target_author).Update("name", update_author.Name)
		messages.GetMessageJSON(c, http.StatusOK, &target_author)
	}
}
