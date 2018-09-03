package main

import (
	"github.com/gin-gonic/gin"
	"booksystem/pkg/api"
)

func main(){
	router := gin.Default()
	api.RunHTTPServer(router)
	router.Run()
}
