package api

import (
	"booksystem/pkg/publisher"
	"github.com/gin-gonic/gin"
)

func GetPublisherAPI(engine *gin.Engine) {
	engine.GET("/api/v1alpha1/publishers", publisher.GetPublisherList)
	engine.POST("/api/v1alpha1/publishers", publisher.CreatePublisher)
	engine.GET("/api/v1alpha1/publishers/:id", publisher.GetPublisherByID)
	engine.PUT("/api/v1alpha1/publishers/:id", publisher.UpdatePublisherByID)
	engine.DELETE("/api/v1alpha1/publishers/:id", publisher.DeletePublisherByID)
}
