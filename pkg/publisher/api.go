package publisher

import (
	"booksystem/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetPublisherByID(ctx *gin.Context) {
	publisherID := ctx.Param("id")

	publisher, err := GetPublisherInfo(publisherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, publisher)
}

func CreatePublisher(ctx *gin.Context) {
	publisher := database.Publisher{}

	err := ctx.BindJSON(&publisher)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	result, err := CreateORUpdatePublisher(publisher, Create, strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetPublisherList(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	start, err := strconv.Atoi(ctx.DefaultQuery("start", "0"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	publishers, err := GetPublishers(limit, start)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, publishers)
}

func DeletePublisherByID(ctx *gin.Context) {
	publisherID := ctx.Param("id")

	err := DeletePublisher(publisherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func UpdatePublisherByID(ctx *gin.Context) {
	publisher := database.Publisher{}

	err := ctx.BindJSON(&publisher)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	publisherID := ctx.Param("id")

	result, err := CreateORUpdatePublisher(publisher, Update, publisherID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
