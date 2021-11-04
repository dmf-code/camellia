package assembler

import (
	"camellia/internal/admin/interfaces/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserToDTO(c *gin.Context) (dto.UserDTO, error) {
	var obj dto.UserDTO
	if err := c.ShouldBind(&obj); err != nil {
		fmt.Println(err)
		return dto.UserDTO{}, err
	}
	return obj, nil
}
