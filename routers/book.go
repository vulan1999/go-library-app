package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/api"
)

// BooksRoute godoc
// @Tag Book
// @Router /books [get]
func BookRouter(r *gin.Engine) {
	book_group := r.Group("/books")
	{
		book_group.GET("/:id", api.GetBookById)
	}
}
