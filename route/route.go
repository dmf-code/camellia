package route

import (
	"camellia/internal/docker/application/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (e *gin.Engine, err error) {
	r := gin.Default()

	container := r.Group("/container")
	{
		containerController := controller.ContainerController{}
		container.GET("list", containerController.List)
		container.POST("new", containerController.NewContainer)
		container.GET("stop", containerController.Stop)
		container.GET("start", containerController.Start)
		container.GET("remove", containerController.Remove)
	}

	image := r.Group("/image")
	{
		imageController := controller.ImageController{}
		image.POST("/build", imageController.Build)
	}

	return r, nil
}
