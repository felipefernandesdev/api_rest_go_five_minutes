package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	router.GET("/albuns", controller.GetAlbuns)

	router.Run("localhost:8080")
}
