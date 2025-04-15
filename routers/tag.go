package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/api"
)

func TagRouter(r *gin.Engine) {
	tag_group := r.Group("/tags")
	{
		tag_group.GET("/", api.GetAllTags)
		tag_group.GET("/:id", api.GetTagById)
	}
}
