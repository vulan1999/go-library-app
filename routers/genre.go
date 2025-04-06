package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/controllers"
)

func GenreRouter(r *gin.Engine) {
	genre_group := r.Group("/genre")
	{
		genre_group.GET("/", controllers.GetAllGenre)
		genre_group.GET("/:id", controllers.GetGenreById)
	}
}
