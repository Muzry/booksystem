package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunHTTPServer(engine *gin.Engine) {
	engine.GET("/api/v1alpha1/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"appName": "BookSystem", "version": "0.0.1"})
	})
	GetBookAPI(engine)
	GetPublisherAPI(engine)
}
