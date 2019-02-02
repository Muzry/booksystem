package book

import (
	"booksystem/pkg/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("id")

	book, err := GetBookInfo(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {

	book := database.Book{}

	err := ctx.BindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	result, err := CreateORUpdateBook(book, Create, strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetBookList(ctx *gin.Context) {
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

	books, err := GetBooks(limit, start)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func DeleteBookByID(ctx *gin.Context) {
	bookID := ctx.Param("id")

	err := DeleteBook(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func UpdateBookByID(ctx *gin.Context) {
	book := database.Book{}

	err := ctx.BindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	bookID := ctx.Param("id")

	result, err := CreateORUpdateBook(book, Update, bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
