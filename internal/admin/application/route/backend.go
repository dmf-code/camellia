package route

import "github.com/gin-gonic/gin"

func BackendGroup(r *gin.RouterGroup) {
	r.POST("/role/create")
}
