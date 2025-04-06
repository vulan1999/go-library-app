package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vulan1999/go-library-app/config"
	"github.com/vulan1999/go-library-app/routers"
)

func init() {
	config.LoadEnviromentVariables()
	config.ConnectToDatabase()
}

func main() {
	fmt.Println("Ready to Go")
	r := gin.Default()
	routers.AuthorRouter(r)
	routers.GenreRouter(r)
	r.Run("localhost:8080")
}
