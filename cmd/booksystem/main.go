package main

import (
	"booksystem/pkg/api"
	"booksystem/pkg/database"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowMethods = append(config.AllowMethods, "DELETE")
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))
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
