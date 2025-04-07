package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/api"
)

func AuthorRouter(r *gin.Engine) {
	author_group := r.Group("/authors")
	{
		author_group.GET("/", api.GetAllAuthor)
		author_group.GET("/:id", api.GetAuthorById)
	}
}
