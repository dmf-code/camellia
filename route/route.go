package route

import (
	"camellia/route/group"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (e *gin.Engine, err error) {
	r := gin.Default()
	docker := r.Group("/docker")
	{
		docker.GET("/ps", group.Ps)
		docker.POST("/run", group.Run)
		docker.POST("/build", group.ImageBuild)
	}
	return r, nil
}
