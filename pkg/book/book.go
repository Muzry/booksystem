package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBookByID(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "get book by ID.")
}

func CreateBook(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "create book.")
}

func GetBookList(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "get book list.")
}

func DeleteBookByID(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "delete book by ID.")
}

func UpdateBookByID(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "update book by ID.")
}