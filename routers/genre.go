package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/api"
)

func GenreRouter(r *gin.Engine) {
	genre_group := r.Group("/genres")
	{
		genre_group.GET("/", api.GetAllGenre)
		genre_group.GET("/:id", api.GetGenreById)
	}
}
