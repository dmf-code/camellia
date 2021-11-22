package route

import (
	"camellia/internal/admin/application/service"
	"github.com/gin-gonic/gin"
)

func BackendGroup(r *gin.RouterGroup) {
	roleApp := service.NewRoleAppService()
	r.POST("/role/create", roleApp.Create)

	userApp := service.NewUserAppService()
	r.GET("/user/login", userApp.Login)
	r.POST("/user/register", userApp.Register)
}
