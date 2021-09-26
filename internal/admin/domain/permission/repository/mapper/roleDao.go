package mapper

import (
	"camellia/internal/admin/domain/user/repository/po"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func (dao RoleDao) Insert(userPO po.UserPO) {
	dao.db.Create(&userPO)
}
