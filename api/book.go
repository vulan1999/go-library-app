package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/models"
	"github.com/vulan1999/go-library-app/utils/messages"
	queryerrors "github.com/vulan1999/go-library-app/utils/query_errors"
	"gorm.io/gorm"
)

// GetBook godoc
//
//	@Summary	Get all books
//	@ID			get-all-books
//	@Tags		Book
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int	false	"List page"
//	@Param		limit	query		int	false	"List limit"
//	@Success	200		{object}	messages.Response{data=[]models.BookResponse}
//	@Failure	502		{object}	messages.Response{data=nil}
//	@Router		/books/ [get]
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	page, page_err := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, limit_err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page_err != nil || limit_err != nil {
		messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
		log.Panic(page_err)
		log.Panic(limit_err)
		return
	}

	offset := 0
	if page > 1 {
		offset = limit * (page - 1)
	}

	result := config.Db.
		Limit(limit).
		Offset(offset).
		Preload("OriginalBook", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id", "Title")
		}).
		Find(&books)

	if result.Error != nil {
		messages.GetMessageJSON(c, http.StatusInternalServerError, nil)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &books)
	}
}

// GetBook godoc
//
//	@Summary	Get book by id
//	@ID			get-book-by-id
//	@Tags		Book
//	@Accept		json
//	@Param		id	path	int	true	"Book ID"
//	@Produce	json
//	@Success	200	{object}	messages.Response{data=models.BookResponse}
//	@Failure	502	{object}	messages.Response{data=nil}
//	@Router		/books/{id} [get]
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
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id", "Description")
		}).
		First(&target, id)

	if result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, result.Error)
		log.Panic(result.Error)
	} else {
		messages.GetMessageJSON(c, http.StatusOK, &target)
	}
}

// @Summary	Create Book
// @Tags		Book
// @ID			create-book
// @Accept		json
// @Produce	json
// @Param		body	body		models.BookCreateRequest	true	"Book detail to create"
// @Success	201		{object}	messages.Response{data=models.BookResponse}
// @Failure	500		{object}	messages.Response{data=nil}
// @Router		/books [post]
func CreateBook(c *gin.Context) {
	var body models.BookCreateRequest

	if err := c.BindJSON(&body); err != nil {
		messages.GetMessageJSON(c, http.StatusBadRequest, nil)
		return
	}

	newBook := models.Book{
		Title:          body.Title,
		AuthorId:       body.AuthorId,
		LanguageId:     body.LanguageId,
		OriginalBookId: body.OriginalBookId,
	}

	create_result := config.Db.Create(&newBook)

	if create_result.Error != nil {
		queryerrors.GetQueryErrorMessage(c, create_result.Error)
		return
	}
	messages.GetMessageJSON(c, http.StatusCreated, &newBook)
}
