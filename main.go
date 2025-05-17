package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/docs"
	"github.com/vulan1999/go-library-app/routers"
)

//	@title			Babel Book Library API
//	@version		1.0
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath		/
func init() {
	config.LoadEnviromentVariables()
	config.ConnectToDatabase()
}

func main() {
	fmt.Println("Ready to Go")
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	routers.AuthorRouter(r)
	routers.GenreRouter(r)
	routers.TagRouter(r)
	routers.BookRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8080")
}
