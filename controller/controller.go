package controller

import (
	"net/http"
	"rest-go/models"

	"github.com/gin-gonic/gin"
)

var albuns = []models.Album{
	{ID: "1", Title: "Equalize", Artist: "Pitty", Price: 50.00},
	{ID: "2", Title: "Planeta de Cores", Artist: "Forrozão Tropicallya", Price: 10.00},
}

func GetAlbuns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albuns)
}
