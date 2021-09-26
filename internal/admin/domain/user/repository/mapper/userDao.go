package mapper

import (
	"camellia/internal/admin/domain/user/repository/po"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func (dao UserDao) Insert(userPO po.UserPO) {
	dao.db.Create(&userPO)
}
