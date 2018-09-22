package publisher

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPublisherByID(ctx *gin.Context){
	publisherID := ctx.Query("publisherid")
	ctx.JSON(http.StatusOK, "get Publisher by ID.")
}

func CreatePublisher(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "create Publisher.")
}

func GetPublisherList(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "get Publisher list.")
}

func DeletePublisherByID(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "delete Publisher by ID.")
}

func UpdatePublisherByID(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "update Publisher by ID.")
}