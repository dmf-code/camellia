package route

import (
	"camellia/internal/admin/application/service"
	"github.com/gin-gonic/gin"
)

func BackendGroup(r *gin.RouterGroup) {
	roleApp := service.RoleApplicationService{}
	r.POST("/role/create", roleApp.Create)
}
