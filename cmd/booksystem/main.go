package main

import (
	"booksystem/pkg/api"
	"booksystem/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RunHTTPServer(router)
	engine, err := database.GetConnection()
	defer engine.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = database.InitTable(engine)

	if err != nil {
		fmt.Println(err)
		return
	}

	router.Run()
}
