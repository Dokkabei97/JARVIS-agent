package router

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	router := gin.Default()

	docker := router.Group("/api/v1/docker")
	{
		docker.GET("", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "docker",
			})
		})
	}

	kubernetes := router.Group("/api/v1/kuber")
	{
		kubernetes.GET("", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "kubernetes",
			})
		})
	}

	return router
}
