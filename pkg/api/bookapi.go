package api

import (
	"github.com/gin-gonic/gin"
	"booksystem/pkg/book"
)

func GetBookAPI(engine *gin.Engine){
	engine.GET("/api/v1alpha1/books", book.GetBookList)
	engine.POST("/api/v1alpha1/books", book.CreateBook)
	engine.GET("/api/v1alpha1/books/:id", book.GetBookByID)
	engine.PUT("/api/v1alpha1/books/:id", book.UpdateBookByID)
	engine.DELETE("/api/v1alpha1/books/:id", book.DeleteBookByID)
}